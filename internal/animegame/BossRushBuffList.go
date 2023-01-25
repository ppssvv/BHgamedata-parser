package animegame

import (
	"encoding/json"
)

type BossRushBuffList struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []BossRushBuffListEntry
}

func (s *BossRushBuffList) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type BossRushBuffListEntry struct {
	ID    int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Unk1  int32

	Description Hash
	Name        Hash
	AbilityName StrWithPrefix16
	IconPath    StrWithPrefix16
}
