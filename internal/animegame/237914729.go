package animegame

import "encoding/json"

// 237914729

type Data_237914729 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_237914729_Entry
}

func (s *Data_237914729) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_237914729_Entry struct {
	ID    int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Unk1  int16
	Addr5 uint32 `json:"-"`
	Unk2  int32
	Addr6 uint32 `json:"-"`

	Hash1   Hash
	Hash2   Hash
	Hash3   Hash
	Hash4   Hash
	UnkStr1 StrWithPrefix16
	UnkStr2 StrWithPrefix16
}
