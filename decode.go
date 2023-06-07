package dataparse

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"dataparse/internal/mt"
)

// DecodeFile is a small wrapper around Decode to work with files
func DecodeFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Decode(f)
}

func Decode(f io.ReadSeeker) ([]byte, error) {
	var seed int32
	if err := binary.Read(f, binary.LittleEndian, &seed); err != nil {
		return nil, err
	}

	var data [0x80]byte
	if err := binary.Read(f, binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	xorpad := [0x200]byte{}

	mt := mt.New(seed)

	// 512 (0x200) bytes or 128 (0x80) ints
	for i := 0; i < 512; i += 4 {
		temp := bytes.Buffer{}

		if err := binary.Write(&temp, binary.LittleEndian, int32(mt.Int31())); err != nil {
			return nil, err
		}

		seedBytes := temp.Bytes()

		xorpad[i] = seedBytes[0] ^ data[i%0x80]
		xorpad[i+1] = seedBytes[1] ^ data[(i+7)%0x80]
		xorpad[i+2] = seedBytes[2] ^ data[(i+13)%0x80]
		xorpad[i+3] = seedBytes[3] ^ data[(i+11)%0x80]
	}

	xorpadIndex := mt.Int31()
	pos := int(xorpadIndex) - 0x84

	size, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	result := make([]byte, size)
	f.Seek(0, io.SeekStart)
	if err := binary.Read(f, binary.LittleEndian, &result); err != nil {
		return nil, err
	}

	for i := 0; i < len(result); i++ {
		result[i] ^= xorpad[(i+pos)%0x200]
	}

	return result[0x84:], nil
}
