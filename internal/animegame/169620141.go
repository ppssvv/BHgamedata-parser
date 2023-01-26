package animegame

import "encoding/json"

// 169620141

type Data_169620141 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_169620141_Entry
}

func (s *Data_169620141) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_169620141_Entry struct {
	ID    int32
	Addr1 uint32 `json:"-"`
	Hash1 Hash
}
