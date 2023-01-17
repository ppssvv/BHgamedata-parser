package animegame

import (
	"dataparse/internal/binreader"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func mustOpenFile(f string) *os.File {
	fd, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	return fd
}

func readHash(reader *binreader.Unpacker) int32 {
	hasHash := reader.Byte()
	if hasHash > 0 {
		if hasHash > 1 {
			panic("hi3")
		}
		return reader.Int32()
	}

	return 0
}

func GetAsset(f string) Asset {
	f = filepath.Base(f)
	id, _, _ := strings.Cut(f, "_")
	a, ok := AssetName[id]
	if ok {
		return a
	}

	return Asset{
		Name: f,
		Parser: func(f string) ([]byte, error) {
			return nil, fmt.Errorf("not implemeted")
		},
	}
}

func arrayStr(reader *binreader.Unpacker, count uint32) []string {
	result := []string{}

	for i := uint32(0); i < count; i++ {
		result = append(result, reader.StringWithUint16Prefix())
	}

	return result
}
