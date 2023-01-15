package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type DormEvent []DormEventEntry

type DormEventEntry struct {
	Hash             int32
	Type             string
	Avatar           int8
	Wait             float32
	WaitRandomAdd    float32
	Emotion          string
	TalkTxt          string
	TalkToAvatar     int8
	Unknown          int8
	FaceAnimType     string
	TriggerAction    string
	TriggerSubAction string
}

func ProcessDormEvent(f string) error {
	obj := NewDormEvent(f)
	result, err := obj.JSON()
	if err != nil {
		return fmt.Errorf("can't marshal DormEvent: %w", err)
	}

	s, err := resultDir("dormEvent")
	if err != nil {
		return fmt.Errorf("can't create output dir [%s]: %w", s, err)
	}

	return os.WriteFile(
		filepath.Join(s, fmt.Sprintf("%s.json", GetAsset(f).Name)),
		result,
		os.ModePerm,
	)
}

func (t *DormEvent) JSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func NewDormEvent(name string) DormEvent {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(name))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 4) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []DormEventEntry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newDormEventEntry(reader))
	}

	return result
}

func newDormEventEntry(reader *binreader.Unpacker) DormEventEntry {
	result := DormEventEntry{}

	// - Header

	result.Hash = reader.Int32()

	reader.Skip(4) // addrto

	result.Avatar = reader.Int8()
	result.Wait = reader.Float32()
	result.WaitRandomAdd = reader.Float32()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.TalkToAvatar = reader.Int8()
	result.Unknown = reader.Int8()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	// - Body

	result.Type = reader.StringWithUint16Prefix()
	result.Emotion = reader.StringWithUint16Prefix()
	result.TalkTxt = reader.StringWithUint16Prefix()
	result.FaceAnimType = reader.StringWithUint16Prefix()
	result.TriggerAction = reader.StringWithUint16Prefix()
	result.TriggerSubAction = reader.StringWithUint16Prefix()

	return result
}
