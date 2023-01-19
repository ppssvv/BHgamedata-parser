package animegame

import (
	"dataparse/internal/binreader"
	"os"
)

type Hash struct {
	Hash int32
}

func readHash(reader *binreader.Unpacker) Hash {
	hasHash := reader.Byte()
	if hasHash > 0 {
		if hasHash > 1 {
			panic("hi3")
		}
		return Hash{reader.Int32()}
	}

	return Hash{0}
}

func mustOpenFile(f string) *os.File {
	fd, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	return fd
}
