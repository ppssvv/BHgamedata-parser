package dump

import (
	"encoding/binary"
	"strings"

	bin "github.com/streamingfast/binary"
)

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
