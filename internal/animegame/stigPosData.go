package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
	"encoding/json"
)

type StigPosData []StigPosDataEntry

// ProcessDormEvent provides a unified interface for batch processing
func ProcessStigPosData(f string) ([]byte, error) {
	return json.MarshalIndent(NewStigPosData(f), "", "  ")
}

type StigPosDataEntry struct {
	Name           string
	RootHeight     float32
	RootWidth      float32
	HasBack        int8
	BackPosX       float32
	BackPosY       float32
	BackHeight     float32
	BackWidth      float32
	BackPath       string
	ImageHeight    float32
	ImageWidth     float32
	ImagePath      string
	HasFront       int8
	FrontPosX      float32
	FrontPosY      float32
	FrontHeight    float32
	FrontWidth     float32
	FrontRotationX float32
	FrontRotationY float32
	FrontRotationZ float32
	FrontPath      string
}

func NewStigPosData(f string) StigPosData {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	// names := arrayStr(reader, uint32(entryCount))
	// list of names
	for i := 0; i < int(entryCount); i++ {
		reader.Skip(int64(reader.Uint16()))
	}

	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := make([]StigPosDataEntry, 0, entryCount)

	for i := 0; i < int(entryCount); i++ {
		tmp := newStigPosDataEntry(reader)
		// tmp.Name = names[i]
		result = append(result, tmp)
	}

	return result
}

func newStigPosDataEntry(reader *binreader.Unpacker) StigPosDataEntry {
	result := StigPosDataEntry{}

	reader.Skip(4) // addrto

	result.RootHeight = reader.Float32()
	result.RootWidth = reader.Float32()

	result.HasBack = reader.Int8()

	result.BackPosX = reader.Float32()
	result.BackPosY = reader.Float32()
	result.BackHeight = reader.Float32()
	result.BackWidth = reader.Float32()

	reader.Skip(4)

	result.ImageHeight = reader.Float32()
	result.ImageWidth = reader.Float32()

	reader.Skip(4)

	result.HasFront = reader.Int8()

	result.FrontPosX = reader.Float32()
	result.FrontPosY = reader.Float32()
	result.FrontHeight = reader.Float32()
	result.FrontWidth = reader.Float32()
	result.FrontRotationX = reader.Float32()
	result.FrontRotationY = reader.Float32()
	result.FrontRotationZ = reader.Float32()

	reader.Skip(4)

	result.Name = reader.StringWithUint16Prefix()
	result.BackPath = reader.StringWithUint16Prefix()
	result.ImagePath = reader.StringWithUint16Prefix()
	result.FrontPath = reader.StringWithUint16Prefix()

	return result
}
