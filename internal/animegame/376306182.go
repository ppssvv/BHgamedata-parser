package animegame

import "encoding/json"

// 376306182

type OWEndlessBattleConfig struct {
	Filesize   uint32                `json:"-"`
	EntryCount uint32                `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []EndlessBattleConfig `json:"-"`
	Junk2      []uint32              `json:"-"`
	Data       []OWEndlessBattleConfig_Entry
}

func (s *OWEndlessBattleConfig) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type OWEndlessBattleConfig_Entry struct {
	EndlessBattleConfig
	Addr1             uint32 `json:"-"`
	Addr2             uint32 `json:"-"`
	Addr3             uint32 `json:"-"`
	MapID             int32
	AreaScoreConfigID int32

	EnvironmentIDList ArrInt32
	AreaIDList        ArrInt32
	BossAreaIDList    ArrInt32
}

type EndlessBattleConfig struct {
	BattleConfigID    int32
	EndlessConfigType int8
}
