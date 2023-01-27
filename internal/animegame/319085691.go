package animegame

import "encoding/json"

// 319085691

type Data_319085691 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint64 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_319085691_Entry
}

func (s *Data_319085691) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_319085691_Entry struct {
	ID   int64
	Unk1 int32
	Unk2 int32
}
