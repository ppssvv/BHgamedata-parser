package dump

import (
	"encoding/binary"
	"encoding/json"
	"strings"

	bin "github.com/streamingfast/binary"
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

func (s *StructReader[K, D]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Data)
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

func (a *AddrTo[C]) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Data)
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type RemoteTimeSpan struct {
	TotalSeconds int32
}

type RemoteTime struct {
	TimeStampForBakedReader uint32
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
