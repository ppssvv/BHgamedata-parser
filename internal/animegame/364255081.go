package animegame

import "encoding/json"

// 364255081

type Data_364255081 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_364255081_Entry
}

func (s *Data_364255081) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_364255081_Entry struct {
	ID    int32
	Unk1  int32
	Addr1 uint32 `json:"-"`
	Arr1  ArrInt32
}
