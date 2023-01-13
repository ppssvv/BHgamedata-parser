package animegame

import (
	"dataparse/internal/ksygen"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TextMap []TextMapEntry

type TextMapEntry struct {
	Hash uint32
	Text string
}

func ProcessTextMap(f string) error {
	obj, err := NewTextMap(f)
	if err != nil {
		return err
	}
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

func (t *TextMap) fillFrom(source *ksygen.TextMap) error {
	texts, err := source.Textlist()
	if err != nil {
		return err
	}

	for _, entry := range texts {
		// temp["key"] = strconv.FormatBool(entry.Key == 1)
		temp := TextMapEntry{}
		temp.Hash = entry.Hash
		temp.Text = entry.Text

		*t = append(*t, temp)
	}

	return nil
}

func (t *TextMap) JSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func NewTextMap(name string) (*TextMap, error) {
	stream, err := getStream(name)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	var result = new(TextMap)
	bin := ksygen.NewTextMap()
	if err := bin.Read(stream, nil, bin); err != nil {
		return nil, fmt.Errorf("can't parse textMap: %w", err)
	}

	if err := result.fillFrom(bin); err != nil {
		return nil, fmt.Errorf("can't parse textMap entries: %w", err)
	}

	return result, nil
}
