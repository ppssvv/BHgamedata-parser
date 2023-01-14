package animegame

import (
	"dataparse/internal/ksygen"
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
	obj, err := NewDormEvent(f)
	if err != nil {
		return err
	}
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

func (t *DormEvent) fillFrom(source *ksygen.DormEvent) error {
	entries, err := source.Entries()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		temp := DormEventEntry{
			Hash:             entry.Header.Hash,
			Type:             entry.Data.EventType.Value,
			Avatar:           entry.Header.Unk1,
			Wait:             entry.Header.Unk2,
			WaitRandomAdd:    entry.Header.Unk3,
			Emotion:          entry.Data.Emotion.Value,
			TalkTxt:          entry.Data.Unk1.Value,
			TalkToAvatar:     entry.Header.Unk4,
			Unknown:          entry.Header.Unk5,
			FaceAnimType:     entry.Data.Unk2.Value,
			TriggerAction:    entry.Data.Unk3.Value,
			TriggerSubAction: entry.Data.Unk4.Value,
		}

		*t = append(*t, temp)
	}

	return nil
}

func (t *DormEvent) JSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func NewDormEvent(name string) (*DormEvent, error) {
	stream, err := getStream(name)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	var result = new(DormEvent)
	bin := ksygen.NewDormEvent()
	if err := bin.Read(stream, nil, bin); err != nil {
		return nil, fmt.Errorf("can't parse textMap: %w", err)
	}

	if err := result.fillFrom(bin); err != nil {
		return nil, fmt.Errorf("can't parse textMap entries: %w", err)
	}

	return result, nil
}
