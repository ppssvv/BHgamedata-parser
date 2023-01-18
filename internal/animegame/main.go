package animegame

import (
	"dataparse/internal/animegame/animetype"
	"encoding/json"
)

// ProcessTextMap provides a unified interface for batch processing
func ProcessTextMap(f string) ([]byte, error) {
	return json.MarshalIndent(animetype.NewTextMap(f), "", "  ")
}

// ProcessDormEvent provides a unified interface for batch processing
func ProcessStigPosData(f string) ([]byte, error) {
	return json.MarshalIndent(animetype.NewStigPosData(f), "", "  ")
}

// ProcessDialogueData provides a unified interface for batch processing
func ProcessOWActivityBossData(f string) ([]byte, error) {
	return json.MarshalIndent(animetype.NewOWActivityBossData(f), "", "  ")
}

// ProcessDormFurniture provides a unified interface for batch processing
func ProcessDormFurniture(f string) ([]byte, error) {
	return json.MarshalIndent(animetype.NewDormFurniture(f), "", "  ")
}

// ProcessDormEvent provides a unified interface for batch processing
func ProcessDormEvent(f string) ([]byte, error) {
	return json.MarshalIndent(animetype.NewDormEvent(f), "", "  ")
}

// ProcessDialogueData provides a unified interface for batch processing
func ProcessDialogueData(f string) ([]byte, error) {
	return json.MarshalIndent(animetype.NewDialogueData(f), "", "  ")
}

// ProcessDormEvent provides a unified interface for batch processing
func ProcessWorldMap(f string) ([]byte, error) {
	return json.MarshalIndent(animetype.NewWorldMap(f), "", "  ")
}
