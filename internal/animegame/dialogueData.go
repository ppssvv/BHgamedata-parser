package animegame

import (
	"dataparse/internal/ksygen"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ProcessDialogueData(f string) error {
	obj, err := NewDialogueData(f)
	if err != nil {
		return err
	}
	result, err := obj.JSON()
	if err != nil {
		return fmt.Errorf("can't marshal DialogueData: %w", err)
	}

	s, err := resultDir("dialogueData")
	if err != nil {
		return fmt.Errorf("can't create output dir [%s]: %w", s, err)
	}

	return os.WriteFile(
		filepath.Join(s, fmt.Sprintf("%s.json", GetAsset(f).Name)),
		result,
		os.ModePerm,
	)
}

type DialogueData []DialogueEntry

type DialogueEntry struct {
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

func (t *DialogueData) fillFrom(source *ksygen.DialogueData) error {
	entries, err := source.Entries()
	if err != nil {
		return fmt.Errorf("can't parse DialogueData: %w", err)
	}

	for _, entry := range entries {
		temp := DialogueEntry{}

		temp.ID = entry.Header.EntryId
		temp.PreDialogueIDList = entry.Data.PostDialogue.Entries
		temp.JumpID = entry.Header.Unk1
		temp.LeafIDList = entry.Data.PreDialogue.Entries
		temp.DialogueType = entry.Header.DialogueType
		temp.InputID = entry.Header.Unk2
		temp.CgRawPos = entry.Header.Unk3
		temp.AvatarName = entry.Data.Avatar.Value
		temp.AvatarID = entry.Header.AvatarId
		temp.DressID = entry.Header.DressId
		temp.AvatarViceKey = entry.Header.AvatarViceKey
		temp.ScreenSide = entry.Header.ScreenSide
		temp.FaceID = entry.Header.FaceId
		temp.Face = entry.Data.Face.Value
		temp.FaceType = entry.Data.FaceAnimation.Value
		temp.FaceEffect = entry.Data.FaceEffect.Value
		temp.AnimationID = entry.Header.AnimationId
		temp.Distortion = entry.Header.Distortion
		temp.Transparency = entry.Header.Transparency
		temp.BGMCover = entry.Data.Unknown6.Value
		temp.BGMVolumeControlList = convertStrArray(entry.Data.Unknown7)
		temp.CgID = entry.Data.Unknown8.Value
		temp.ImageID = entry.Header.Unk5
		temp.Content = convertContent(entry.Data.Content)
		temp.Background = entry.Data.Unknown9.Value
		temp.BackgroundCG = entry.Data.Unknown10.Value
		temp.BackgroundEffect = entry.Header.BackgroundEffect
		temp.EnterEffect = entry.Header.EnterEffect
		temp.AudioID = entry.Data.AudioId.Value
		temp.AudioDelay = entry.Header.AudioDelay
		temp.LipMotion = entry.Data.LipMotion.Value
		temp.FaceVersion = entry.Header.Unk7
		temp.ScreenEffect = entry.Data.Unknown11.Value
		temp.Unknown1 = entry.Data.Unknown12.Value
		temp.Unknown2 = convertStrArray(entry.Data.Unknown13)
		temp.Unknown3 = convertStrArray(entry.Data.Unknown14)
		temp.Unknown4 = entry.Header.Unk8
		temp.Unknown5 = entry.Header.Unk9
		temp.PostDialogueIDList = entry.Data.Unknown15.Entries
		temp.TalkerName = entry.Data.Unknown16.Value
		temp.QuestionContent = entry.Data.Unknown17.Value

		*t = append(*t, temp)
	}

	return nil
}

func convertStrArray(source *ksygen.DialogueData_StrArr) []string {
	result := []string{}

	for _, e := range source.Entries {
		result = append(result, e.Value)
	}

	return result
}

func convertContent(source *ksygen.DialogueData_DialogueContent) []DialogueContent {
	result := []DialogueContent{}

	for _, e := range source.Dialogues {
		result = append(result, DialogueContent{Text: e.Value, Duration: e.Duration})
	}

	return result
}

func (t *DialogueData) JSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func NewDialogueData(name string) (*DialogueData, error) {
	stream, err := getStream(name)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	var result = new(DialogueData)
	bin := ksygen.NewDialogueData()
	if err := bin.Read(stream, nil, bin); err != nil {
		return nil, fmt.Errorf("can't parse dialogueData: %w", err)
	}

	if err := result.fillFrom(bin); err != nil {
		return nil, fmt.Errorf("can't parse dialogueData entries: %w", err)
	}

	return result, nil
}
