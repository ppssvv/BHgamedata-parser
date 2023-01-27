package animegame

import "encoding/json"

// 226923750

type Data_226923750 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_226923750_Entry
}

func (s *Data_226923750) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_226923750_Entry struct {
	ID    int32
	Unk1  int32
	Addr1 uint32 `json:"-"`
	Unk2  int32
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`

	UnkStr1 StrWithPrefix16
	Hash1   Hash
	UnkStr2 StrWithPrefix16
}
