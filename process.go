package dataparse

import (
	"dataparse/internal/animegame"
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

func GetTestData(f string) *bin.Decoder {
	if filepath.IsAbs(f) {
		return readFile(f)
	} else {
		return readFile(filepath.Join("testdata", f))
	}
}

func ProcessStruct(f string, obj animegame.GameStruct) ([]byte, error) {
	if err := GetTestData(f).Decode(obj); err != nil {
		return nil, err
	}

	return obj.JSON()
}
