package animegame

import (
	"encoding/json"
)

type WorldMap struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []WorldMapEntry
}

func (s *WorldMap) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type WorldMapEntry struct {
	ID      int32
	EntryID int32

	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	EntryCG          int32
	UnlockLevel      int32
	UnlockHintLevel  int32
	ForceShowContent int32
	InfoID           int32
	ShopType         int32

	Addr6 uint32 `json:"-"`
	Addr7 uint32 `json:"-"`

	EntryRewardID   int32
	UnkByte1        byte
	VersionTagMaxLv int32
	VersionTagMinLv int32

	UnkByte2 byte

	FeatureType       ArrInt32
	ContentIDList     ArrInt32
	EntryTitle        Hash
	EntryImagePath    StrWithPrefix16
	EntryTitleImgPath StrWithPrefix16

	EntryDesc            Hash
	EntryDetailImagePath StrWithPrefix16
}
