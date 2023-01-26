package animegame

import "encoding/json"

// 213532048

type Data_213532048 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_213532048_Entry
}

func (s *Data_213532048) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_213532048_Entry struct {
	ID    int32
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	Hash1 Hash
	Hash2 Hash
	Str1  StrWithPrefix16
	Str2  StrWithPrefix16
	Str3  StrWithPrefix16
}
