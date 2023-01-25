package dataparse

import (
	"dataparse/internal/animegame"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
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

	"1435715286": {"DormitoryEventSequence", &animegame.DormEvent{}},
	"606639383":  {"DormitoryFurnitureData", &animegame.DormFurniture{}},

	"612883459": {"StigmataPositionData", &animegame.StigPosData{}},

	"11165326": {"OWActivityBossData", &animegame.OWActivityBossData{}},

	"36243594": {"WorldMap", &animegame.WorldMap{}},

	"53370678": {"BossRushBuffList", &animegame.BossRushBuffList{}},

	"50428269":  {"50428269", &animegame.Data_50428269{}},
	"112154430": {"112154430", &animegame.Data_112154430{}},
	"137850209": {"137850209", &animegame.Data_137850209{}},
	"150008597": {"150008597", &animegame.Data_150008597{}},
}

type Asset struct {
	Name   string
	Parser animegame.GameStruct
}

func GetAsset(f string) Asset {
	f = filepath.Base(f)
	id, _, _ := strings.Cut(f, "_")
	a, ok := AssetName[id]
	if ok {
		return a
	}

	return Asset{
		Name:   f,
		Parser: Placeholder{},
	}
}

func dialogueDataAsset(lang string) Asset {
	return Asset{
		Name:   fmt.Sprintf("dialogueData_%s", lang),
		Parser: &animegame.DialogueData{},
	}
}

func textMapAsset(lang string) Asset {
	return Asset{
		Name:   fmt.Sprintf("textMap_%s", lang),
		Parser: &animegame.TextMap{},
	}
}

type Placeholder struct{}

func (p Placeholder) JSON() ([]byte, error) {
	return nil, errors.New("not implemented")
}
