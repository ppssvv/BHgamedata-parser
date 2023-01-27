package animegame

import "encoding/json"

// 250894170

type Data_250894170 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint16 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_250894170_Entry
}

func (s *Data_250894170) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_250894170_Entry struct {
	ID    int16
	Unk1  [3]int8
	Addr1 uint32 `json:"-"`
	Unk2  int32
	Unk3  int32
	Unk4  int32
	Unk5  int8

	UnkStr1 StrWithPrefix16
}
