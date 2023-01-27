package animegame

import "encoding/json"

// 220083180

type Data_220083180 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_220083180_Entry
}

func (s *Data_220083180) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_220083180_Entry struct {
	ID    int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Unk   [9]int32

	UnkStr1 StrWithPrefix16
	UnkStr2 StrWithPrefix16
}
