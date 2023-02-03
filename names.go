package dataparse

import (
	"dataparse/dump"
	"dataparse/internal/animegame"
	"errors"
	"path/filepath"
	"strings"
)

var NewAssets = map[string]dump.Asset{
	"2578607515": {Name: dump.AssetName["1453"].Name + "_cn", Parser: dump.AssetName["1453"].Parser},
	"2578607537": {Name: dump.AssetName["1453"].Name + "_de", Parser: dump.AssetName["1453"].Parser},
	"2578607577": {Name: dump.AssetName["1453"].Name + "_en", Parser: dump.AssetName["1453"].Parser},
	"2578607612": {Name: dump.AssetName["1453"].Name + "_fr", Parser: dump.AssetName["1453"].Parser},

	// "816421621": {Name: dump.AssetName["508"].Name + "_fr", Parser: dump.AssetName["1453"].Parser},
	// "816421643": {Name: dump.AssetName["508"].Name + "_fr", Parser: dump.AssetName["1453"].Parser},
	// "816421683": {Name: dump.AssetName["508"].Name + "_fr", Parser: dump.AssetName["1453"].Parser},
	// "816421718": {Name: dump.AssetName["508"].Name + "_fr", Parser: dump.AssetName["1453"].Parser},

	"1435715286": dump.AssetName["435"],
	"606639383":  dump.AssetName["440"],
	"612883459":  dump.AssetName["1426"],
	"11165326":   dump.AssetName["1005"],
	// "36243594":   dump.AssetName["975"],
	"53370678":  dump.AssetName["158"],
	"376306182": dump.AssetName["1022"],
	"379497939": dump.AssetName["111"],
}

var AssetName = map[string]Asset{
	// "2578607515": textMapAsset("cn"),
	// "2578607537": textMapAsset("de"),
	// "2578607577": textMapAsset("en"),
	// "2578607612": textMapAsset("fr"),

	// "816421621": dialogueDataAsset("cn"),
	// "816421643": dialogueDataAsset("de"),
	// "816421683": dialogueDataAsset("en"),
	// "816421718": dialogueDataAsset("fr"),

	// "1435715286": {"DormitoryEventSequence", &animegame.DormEvent{}},
	// "606639383":  {"DormitoryFurnitureData", &animegame.DormFurniture{}},

	// "612883459": {"StigmataPositionData", &animegame.StigPosData{}},

	// "11165326": {"OWActivityBossData", &animegame.OWActivityBossData{}},

	// "36243594": {"WorldMap", &animegame.WorldMap{}},

	// "53370678": {"BossRushBuffList", &animegame.BossRushBuffList{}},
	// "376306182": {"OWEndlessBattleConfig", &animegame.OWEndlessBattleConfig{}},
	// "379497939": {"AvatarFragmentData", &animegame.AvatarFragmentData{}},

	"50428269":  {"50428269", &animegame.Data_50428269{}},
	"112154430": {"112154430", &animegame.Data_112154430{}},
	"137850209": {"137850209", &animegame.Data_137850209{}},
	"150008597": {"150008597", &animegame.Data_150008597{}},
	"169620141": {"169620141", &animegame.Data_169620141{}},
	"175952655": {"175952655", &animegame.Data_175952655{}},
	"201380970": {"201380970", &animegame.Data_201380970{}},
	"208715169": {"208715169", &animegame.Data_208715169{}},
	"213532048": {"213532048", &animegame.Data_213532048{}},
	"220083180": {"220083180", &animegame.Data_220083180{}},
	"226923750": {"226923750", &animegame.Data_226923750{}},
	"237914729": {"237914729", &animegame.Data_237914729{}},
	"244665338": {"244665338", &animegame.Data_244665338{}},
	"250894170": {"250894170", &animegame.Data_250894170{}},
	"285372794": {"285372794", &animegame.Data_285372794{}},
	"295065595": {"295065595", &animegame.Data_295065595{}},
	"319085691": {"319085691", &animegame.Data_319085691{}},
	"354247105": {"354247105", &animegame.Data_354247105{}},
	"364255081": {"364255081", &animegame.Data_364255081{}},
}

type Asset struct {
	Name   string
	Parser animegame.GameStruct
}

func GetAsset(f string) dump.Asset {
	f = filepath.Base(f)
	id, _, _ := strings.Cut(f, "_")
	a, ok := NewAssets[id]
	if ok {
		return a
	}

	return dump.Asset{
		Name:   f,
		Parser: &Placeholder{},
	}
}

// func dialogueDataAsset(lang string) Asset {
// 	return Asset{
// 		Name:   fmt.Sprintf("dialogueData_%s", lang),
// 		Parser: &animegame.DialogueData{},
// 	}
// }

// func textMapAsset(lang string) Asset {
// 	return Asset{
// 		Name:   fmt.Sprintf("textMap_%s", lang),
// 		Parser: &animegame.TextMap{},
// 	}
// }

type Placeholder struct{}

func (p *Placeholder) JSON() ([]byte, error) {
	return nil, errors.New("not implemented")
}
