package dump

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	bin "github.com/streamingfast/binary"
	"golang.org/x/exp/slices"
)

type StructReader[K any, D any] struct {
	Filesize    uint32
	EntryCount  uint32
	Keys        []K
	ItemOffsets []uint32

	Data []D
}

func (s *StructReader[K, D]) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	// totalSize := decoder.Remaining()

	s.Filesize, err = decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return
	}

	// disable check because of 2889932192
	// if s.Filesize != uint32(totalSize) {
	// 	return errors.New("size mismatch")
	// }

	s.EntryCount, err = decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return
	}

	s.Keys = make([]K, s.EntryCount)
	for i := uint32(0); i < s.EntryCount; i++ {
		if err := decoder.Decode(&s.Keys[i]); err != nil {
			return err
		}
	}

	s.ItemOffsets = make([]uint32, s.EntryCount)
	for i := uint32(0); i < s.EntryCount; i++ {
		if err := decoder.Decode(&s.ItemOffsets[i]); err != nil {
			return err
		}
	}

	s.Data = make([]D, s.EntryCount)
	for i := uint32(0); i < s.EntryCount; i++ {
		if err := decoder.SetPosition(uint(s.ItemOffsets[i])); err != nil {
			return nil
		}

		if err := decoder.Decode(&s.Data[i]); err != nil {
			return err
		}
	}

	return nil
}

func IsSimpleKey(s string) bool {
	var simpleKeys = []string{
		"int8",
		"uint8",
		"uint16",
		"uint16",
		"int16",
		"uint32",
		"int32",
		"float32",
		"int64",
		"uint64",
		"Hash",
		"StrWithPrefix16",
		"RemoteTime",
	}

	return slices.Contains(simpleKeys, s)
}

func (s *StructReader[K, D]) MarshalJSON() ([]byte, error) {
	if len(s.Keys) == 0 || len(s.Data) == 0 {
		return []byte("{}"), nil
	}

	// all keys should be same, so only check 1
	if IsSimpleKey(reflect.TypeOf(s.Keys[0]).Name()) {
		// for simple keys we can just make a 1-level map
		var tmp = map[string]D{}
		for i := 0; i < int(s.EntryCount); i++ {
			tmp[fmt.Sprintf("%v", s.Keys[i])] = s.Data[i]
		}

		return json.MarshalIndent(tmp, "", "  ")
	}

	// for complex keys we need to make multilevel maps
	var tmp = map[string]interface{}{}
	for i := 0; i < int(s.EntryCount); i++ {
		v := reflect.ValueOf(s.Keys[i])
		if v.Kind() != reflect.Struct {
			panic("wrong struct")
		}

		var keys = []string{}
		for fi := 0; fi < v.NumField(); fi++ {
			x := getString(v.Field(fi))
			keys = append(keys, x)
		}

		multiMap(tmp, keys, s.Data[i])
	}

	return json.MarshalIndent(tmp, "", "  ")
}

// multiMap works thanks to the power of recursion
func multiMap(mapObj map[string]interface{}, keys []string, finalData interface{}) {
	if len(keys) == 1 {
		mapObj[keys[0]] = finalData
		return
	}

	innerMap, ok := mapObj[keys[0]]
	if !ok {
		innerMap = map[string]interface{}{}
	}

	inner := innerMap.(map[string]interface{})

	multiMap(inner, keys[1:], finalData)
	mapObj[keys[0]] = inner
}

func getString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', 8, 64)
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	default:
		return v.String()
	}
}

type AddrTo[C any] struct {
	Addr uint32
	Data C
}

func (a *AddrTo[C]) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	a.Addr, err = decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return
	}

	if a.Addr <= 0 {
		return
	}

	pos := decoder.Position()
	if err = decoder.SetPosition(uint(a.Addr)); err != nil {
		return
	}

	if err = decoder.Decode(&a.Data); err != nil {
		return
	}

	return decoder.SetPosition(pos)
}

func (a AddrTo[C]) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Data)
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type RemoteTimeSpan int32

type RemoteTime uint32

// Hash actually takes 5 bytes: 1 key + 4 hash
type Hash int32

func (h *Hash) UnmarshalBinary(decoder *bin.Decoder) error {
	key, err := decoder.ReadInt8()
	if err != nil {
		return err
	}

	if key == 0 {
		return nil
	}

	hash, err := decoder.ReadInt32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*h = Hash(hash)
	return nil
}

type StrWithPrefix16 string

func (s *StrWithPrefix16) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	len, err := decoder.ReadUint16(binary.LittleEndian)
	if err != nil {
		return err
	}

	str := strings.Builder{}
	for i := uint16(0); i < len; i++ {
		b, err := decoder.ReadByte()
		if err != nil {
			return err
		}

		_ = str.WriteByte(b)
	}

	*s = StrWithPrefix16(str.String())
	return nil
}

type FileInfo struct {
	Name   Hash
	Offset int32 // from base
	Length int32
}

type CombinedFileInfo struct {
	AssetData []byte
	AssetInfo Asset
}

type CombinedBytesReader struct {
	Filesize   uint32
	BaseOffset uint32
	EntryCount uint16

	FilesInfo []FileInfo
	RawData   []byte
}

func (s *CombinedBytesReader) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	totalSize := decoder.Remaining()

	s.Filesize, err = decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return
	}

	if s.Filesize != uint32(totalSize) {
		return errors.New("size mismatch")
	}

	s.BaseOffset, err = decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return
	}

	s.EntryCount, err = decoder.ReadUint16(binary.LittleEndian)
	if err != nil {
		return
	}

	s.FilesInfo = make([]FileInfo, s.EntryCount)
	for i := uint16(0); i < s.EntryCount; i++ {
		if err := decoder.Decode(&s.FilesInfo[i]); err != nil {
			return err
		}
	}

	s.RawData = decoder.GetData()[s.BaseOffset:]

	return nil
}

func (s *CombinedBytesReader) GetFiles() []CombinedFileInfo {
	result := []CombinedFileInfo{}

	for _, f := range s.FilesInfo {
		result = append(result, CombinedFileInfo{
			AssetData: s.RawData[uint32(f.Offset) : uint32(f.Offset)+uint32(f.Length)],
			AssetInfo: AssetName[BakedHashes[int(f.Name)]],
		})
	}

	return result
}

func (s *CombinedBytesReader) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(s.FilesInfo, "", "  ")
}
