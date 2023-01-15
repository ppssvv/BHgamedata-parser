package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TextMap []TextMapEntry

type TextMapEntry struct {
	Hash int32
	Text string
}

func ProcessTextMap(f string) error {
	obj := NewTextMap(f)
	result, err := obj.JSON()
	if err != nil {
		return fmt.Errorf("can't marshal TextMap: %w", err)
	}

	s, err := resultDir("textMap")
	if err != nil {
		return fmt.Errorf("can't create output dir [%s]: %w", s, err)
	}

	return os.WriteFile(
		filepath.Join(s, fmt.Sprintf("%s.json", GetAsset(f).Name)),
		result,
		os.ModePerm,
	)
}

func (t *TextMap) JSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func NewTextMap(name string) TextMap {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(name))

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
