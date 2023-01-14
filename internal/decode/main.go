package decode

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
)

func Parse(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var seed int32
	if err := binary.Read(f, binary.LittleEndian, &seed); err != nil {
		return nil, err
	}

	var data [0x80]byte
	if err := binary.Read(f, binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	xorpad := [0x200]byte{}

	mt := New(seed)

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

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	result := make([]byte, info.Size())
	f.Seek(0, io.SeekStart)
	if err := binary.Read(f, binary.LittleEndian, &result); err != nil {
		return nil, err
	}

	for i := 0; i < len(result); i++ {
		result[i] ^= xorpad[(i+pos)%0x200]
	}

	return result[0x84:], nil
}
