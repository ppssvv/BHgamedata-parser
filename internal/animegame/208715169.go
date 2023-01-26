package animegame

import "encoding/json"

// 208715169

type Data_208715169 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_208715169_Entry
}

func (s *Data_208715169) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_208715169_Entry struct {
	ID   int32
	Unk1 int32
	Unk2 int32
	Unk3 int32
	Unk4 int32
	Unk5 int32
	Unk6 int32
}
