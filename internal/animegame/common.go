package animegame

import (
	"encoding/binary"
	"strings"

	bin "github.com/streamingfast/binary"
)

type GameStruct interface {
	JSON() ([]byte, error)
}

type Hash struct {
	Hash int32
	Key  int8 `json:",omitempty"`
}

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

	h.Hash = hash
	if key > 1 {
		h.Key = key
	}
	return nil
}

type ArrInt32 []int32

func (arr *ArrInt32) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	count, err := decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*arr = make([]int32, count)
	for i := uint32(0); i < count; i++ {
		(*arr)[i], err = decoder.ReadInt32(binary.LittleEndian)
		if err != nil {
			return err
		}
	}
	return nil
}

type ArrUint32 []uint32

func (arr *ArrUint32) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	count, err := decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*arr = make([]uint32, count)
	for i := uint32(0); i < count; i++ {
		(*arr)[i], err = decoder.ReadUint32(binary.LittleEndian)
		if err != nil {
			return err
		}
	}
	return nil
}

type ArrInt8 []int8

func (arr *ArrInt8) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	count, err := decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*arr = make([]int8, count)
	for i := uint32(0); i < count; i++ {
		(*arr)[i], err = decoder.ReadInt8()
		if err != nil {
			return err
		}
	}
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

type ArrStrWithPrefix16 []StrWithPrefix16

func (arr *ArrStrWithPrefix16) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	count, err := decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*arr = make([]StrWithPrefix16, count)
	for i := uint32(0); i < count; i++ {
		var tmp StrWithPrefix16
		err = tmp.UnmarshalBinary(decoder)
		if err != nil {
			return err
		}
		(*arr)[i] = tmp
	}

	return nil
}
