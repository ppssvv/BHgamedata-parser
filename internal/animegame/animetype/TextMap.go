package animetype

import (
	"dataparse/internal/binreader"
	"encoding/binary"
)

type TextMap []TextMapEntry

type TextMapEntry struct {
	Hash Hash
	Text string
}

func NewTextMap(f string) TextMap {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 5) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []TextMapEntry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newTextMapEntry(reader))
	}

	return result
}

func newTextMapEntry(reader *binreader.Unpacker) TextMapEntry {
	result := TextMapEntry{}

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.Hash = readHash(reader)
	result.Text = reader.StringWithUint16Prefix()

	return result
}