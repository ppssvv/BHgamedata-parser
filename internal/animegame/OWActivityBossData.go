package animegame

import (
	"encoding/json"
)

type OWActivityBossData struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []OWActivityBossDataEntry
}

func (s *OWActivityBossData) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type OWActivityBossDataEntry struct {
	BossId int32

	Addr1 uint32 `json:"-"`

	HardLevel int32
	MonsterHp int32

	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	StaminaCost    int32
	ClueProgress   int32
	ClueTimeLimit  int32
	ClueBuffId     int32
	MutationBuffId int32
	ActivityExp    int32
	BossReward     int32

	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`
	Addr8 uint32 `json:"-"`

	StageId        int32
	MPDamageRatio  float32
	StageMonsterId int32

	MonsterIdList ArrInt32

	ClueEventIdList   ArrInt32
	ClueLocationId    ArrInt32
	ClueTimeLimitTips Hash

	BossName         Hash
	BossDescription  Hash
	BossLocationList ArrInt32
	ImagePath        StrWithPrefix16
}
