package animegame

import "encoding/json"

// 295065595

type Data_295065595 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_295065595_Entry
}

func (s *Data_295065595) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_295065595_Entry struct {
	ID      int32
	Unk1    int32
	Unk2    int32
	Addr1_4 [4]uint32 `json:"-"`
	Unk3    int8
	Unk4    int32
	Unk5    int32
	Unk6    int32
	Unk7    int32
	Unk8    int32
	Addr5   uint32 `json:"-"`
	Unk9    int32

	Str1  StrWithPrefix16
	Str2  StrWithPrefix16
	Str3  StrWithPrefix16
	Str4  StrWithPrefix16
	Hash1 Hash
}
