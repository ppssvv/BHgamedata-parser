package dataparse

import (
	"dataparse/internal/decode"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	bin "github.com/streamingfast/binary"
)

func readFile(f string) *bin.Decoder {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return bin.NewDecoder(data)
}

func checkEncrypted(f string) bool {
	fd, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	var size int32
	if err := binary.Read(fd, binary.LittleEndian, &size); err != nil {
		panic(err)
	}
	defer fd.Close()

	stats, err := fd.Stat()
	if err != nil {
		panic(err)
	}

	return stats.Size() != int64(size)
}

func GetTestData(f string) *bin.Decoder {
	if !filepath.IsAbs(f) {
		f = filepath.Join("testdata", f)
	}

	// before reading all file, check first 4 bytes
	if checkEncrypted(f) {
		fmt.Println("\nfile is encrypted, dec")
		decodedData, err := decode.Parse(f)
		if err != nil {
			panic(err)
		}
		return bin.NewDecoder(decodedData)
	}

	return readFile(f)
}

var ErrNotSupported = errors.New("this file is not supported yet")

func ProcessStructNew(f string, obj any) ([]byte, error) {
	if obj == nil {
		return nil, ErrNotSupported
	}

	if err := GetTestData(f).Decode(obj); err != nil {
		return nil, err
	}

	return json.MarshalIndent(obj, "", "  ")
}
