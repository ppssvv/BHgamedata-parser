package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
	"encoding/json"
)

type DormEvent []DormEventEntry

// ProcessDormEvent provides a unified interface for batch processing
func ProcessDormEvent(f string) ([]byte, error) {
	return json.MarshalIndent(NewDormEvent(f), "", "  ")
}

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

func NewDormEvent(f string) DormEvent {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

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
