package animegame

import "encoding/json"

// 379497939

type AvatarFragmentData struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []AvatarFragmentDataEntry
}

func (s *AvatarFragmentData) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type AvatarFragmentDataEntry struct {
	ID               int32
	Unk1             int32
	Unk2             int32
	Unk3             int32
	MaxLvl           int32
	SellPriceBase    float32
	Unk6             int32
	Unk7             int32
	Unk8             int32
	BaseType         int32
	UnkBytes         [5]int8
	Addr1            uint32 `json:"-"`
	Addr2            uint32 `json:"-"`
	Addr3            uint32 `json:"-"`
	Unk10            int32
	QuantityLimit    int32
	Addr4            uint32 `json:"-"`
	ReturnMaterialID int32
	ReturnNum        int32
	TagType          int8
	Unk14            [3]int8
	Addr5            uint32 `json:"-"`
	UnkByte          int8
	Addr6            uint32 `json:"-"`
	DisplayTitle     Hash

	DisplayDescription Hash
	IconPath           StrWithPrefix16
	ImagePath          StrWithPrefix16
	Arr1               ArrInt32
	Arr2               ArrInt32
	Arr3               ArrInt32
}
