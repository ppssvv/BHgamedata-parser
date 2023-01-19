package dataparse

import (
	"dataparse/internal/animegame"
	"encoding/json"
)

// ProcessTextMap provides a unified interface for batch processing
func ProcessTextMap(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewTextMap(f), "", "  ")
}

// ProcessDormEvent provides a unified interface for batch processing
func ProcessStigPosData(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewStigPosData(f), "", "  ")
}

// ProcessDialogueData provides a unified interface for batch processing
func ProcessOWActivityBossData(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewOWActivityBossData(f), "", "  ")
}

// ProcessDormFurniture provides a unified interface for batch processing
func ProcessDormFurniture(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewDormFurniture(f), "", "  ")
}

// ProcessDormEvent provides a unified interface for batch processing
func ProcessDormEvent(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewDormEvent(f), "", "  ")
}

// ProcessDialogueData provides a unified interface for batch processing
func ProcessDialogueData(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewDialogueData(f), "", "  ")
}

// ProcessDormEvent provides a unified interface for batch processing
func ProcessWorldMap(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewWorldMap(f), "", "  ")
}

func ProcessBossRushBuffList(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewBossRushBuffList(f), "", "  ")
}

func Process_50428269(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewData_50428269(f), "", " ")
}

func Process_112154430(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewData_112154430(f), "", "  ")
}

func Process_137850209(f string) ([]byte, error) {
	return json.MarshalIndent(animegame.NewData_137850209(f), "", "  ")
}
