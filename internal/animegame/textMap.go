package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
	"encoding/json"
)

type TextMap []TextMapEntry

// ProcessTextMap provides a unified interface for batch processing
func ProcessTextMap(f string) ([]byte, error) {
	return json.MarshalIndent(NewTextMap(f), "", "  ")
}

type TextMapEntry struct {
	Hash int32
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
