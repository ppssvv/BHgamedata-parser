package animegame

import "encoding/json"

// 175952655

type Data_175952655 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_175952655_Entry
}

func (s *Data_175952655) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_175952655_Entry struct {
	ID    int32
	Addr1 uint32 `json:"-"`
	Unk1  int32
	Arr1  Data_50428269_ // reusing struct from diff file
}
