package animegame

import (
	"fmt"
)

var AssetName = map[string]Asset{
	"2578607515": textMapAsset("cn"),
	"2578607537": textMapAsset("de"),
	"2578607577": textMapAsset("en"),
	"2578607612": textMapAsset("fr"),

	"816421621": dialogueDataAsset("cn"),
	"816421643": dialogueDataAsset("de"),
	"816421683": dialogueDataAsset("en"),
	"816421718": dialogueDataAsset("fr"),

	"1435715286": {"DormitoryEventSequence", ProcessDormEvent},
	// "606639384":  {"DormitoryFurnitureData", "dormFurn"},

	// "612883459": {"StigmataPositionData", "stigmataPos"},
}

type Asset struct {
	Name   string
	Parser func(f string) ([]byte, error)
}

func dialogueDataAsset(lang string) Asset {
	return Asset{
		Name:   fmt.Sprintf("dialogueData_%s", lang),
		Parser: ProcessDialogueData,
	}
}

func textMapAsset(lang string) Asset {
	return Asset{
		Name:   fmt.Sprintf("textMap_%s", lang),
		Parser: ProcessTextMap,
	}
}
