package animegame

import "encoding/json"

// 244665338

type Data_244665338 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_244665338_Entry
}

func (s *Data_244665338) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_244665338_Entry struct {
	ID    int32
	Unk1  int8
	Addr1 uint32 `json:"-"`

	UnkStr1 StrWithPrefix16
}
