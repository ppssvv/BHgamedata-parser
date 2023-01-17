package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
	"encoding/json"
)

type DialogueData []DialogueDataEntry

// ProcessDialogueData provides a unified interface for batch processing
func ProcessDialogueData(f string) ([]byte, error) {
	return json.MarshalIndent(NewDialogueData(f), "", "  ")
}

type DialogueDataEntry struct {
	ID                   int32
	PreDialogueIDList    []uint32
	JumpID               int32
	LeafIDList           []uint32
	DialogueType         int8
	InputID              int32
	CgRawPos             int32
	AvatarName           string
	AvatarID             int32
	DressID              int32
	AvatarViceKey        int32
	ScreenSide           int8
	FaceID               int8
	Face                 string
	FaceType             string
	FaceEffect           string
	AnimationID          int8
	Distortion           int8
	Transparency         float32
	BGMCover             string
	BGMVolumeControlList []string
	CgID                 string
	ImageID              int32
	Content              []DialogueContent
	Background           string
	BackgroundCG         string
	BackgroundEffect     int8
	EnterEffect          int8
	AudioID              string
	AudioDelay           float32
	LipMotion            string
	FaceVersion          int8
	ScreenEffect         string
	Unknown1             string
	Unknown2             []string
	Unknown3             []string
	Unknown4             int32
	Unknown5             int32
	PostDialogueIDList   []uint32
	TalkerName           string
	QuestionContent      string
}

type DialogueContent struct {
	Text     string
	Duration float32
}

func NewDialogueData(f string) DialogueData {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 4) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []DialogueDataEntry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newDialogueDataEntry(reader))
	}

	return result
}

func newDialogueDataEntry(reader *binreader.Unpacker) DialogueDataEntry {
	result := DialogueDataEntry{}

	// - Header

	result.ID = reader.Int32()

	reader.Skip(4) // addrto

	result.JumpID = reader.Int32()

	reader.Skip(4) // addrto

	result.DialogueType = reader.Int8()
	result.InputID = reader.Int32()
	result.CgRawPos = reader.Int32()

	reader.Skip(4) // addrto

	result.AvatarID = reader.Int32()
	result.DressID = reader.Int32()
	result.AvatarViceKey = reader.Int32()

	result.ScreenSide = reader.Int8()
	result.FaceID = reader.Int8()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.AnimationID = reader.Int8()
	result.Distortion = reader.Int8()
	result.Transparency = reader.Float32()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.ImageID = reader.Int32()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.BackgroundEffect = reader.Int8()
	result.EnterEffect = reader.Int8()

	reader.Skip(4) // addrto

	result.AudioDelay = reader.Float32()

	reader.Skip(4) // addrto

	result.FaceVersion = reader.Int8()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.Unknown4 = reader.Int32()
	result.Unknown5 = reader.Int32()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	// - Body

	result.PreDialogueIDList = reader.ArrayUint32(uint64(reader.Uint32()))
	result.LeafIDList = reader.ArrayUint32(uint64(reader.Uint32()))
	result.AvatarName = reader.StringWithUint16Prefix()
	result.Face = reader.StringWithUint16Prefix()
	result.FaceType = reader.StringWithUint16Prefix()
	result.FaceEffect = reader.StringWithUint16Prefix()
	result.BGMCover = reader.StringWithUint16Prefix()

	result.BGMVolumeControlList = arrayStr(reader, reader.Uint32())

	result.CgID = reader.StringWithUint16Prefix()

	// content
	content := []DialogueContent{}
	contentCount := reader.Uint32()
	for i := 0; i < int(contentCount); i++ {
		content = append(content, newDialogueContent(reader))
	}
	result.Content = content

	result.Background = reader.StringWithUint16Prefix()
	result.BackgroundCG = reader.StringWithUint16Prefix()

	result.AudioID = reader.StringWithUint16Prefix()
	result.LipMotion = reader.StringWithUint16Prefix()

	result.ScreenEffect = reader.StringWithUint16Prefix()
	result.Unknown1 = reader.StringWithUint16Prefix()

	result.Unknown2 = arrayStr(reader, reader.Uint32())
	result.Unknown3 = arrayStr(reader, reader.Uint32())

	result.PostDialogueIDList = reader.ArrayUint32(uint64(reader.Uint32()))
	result.TalkerName = reader.StringWithUint16Prefix()
	result.QuestionContent = reader.StringWithUint16Prefix()

	return result
}

func newDialogueContent(reader *binreader.Unpacker) DialogueContent {
	result := DialogueContent{}

	result.Text = reader.StringWithUint16Prefix()
	result.Duration = reader.Float32()

	return result
}
