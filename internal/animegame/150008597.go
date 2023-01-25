package animegame

import "encoding/json"

// 150008597

type Data_150008597 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_150008597_Entry
}

func (s *Data_150008597) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_150008597_Entry struct {
	ID int32

	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`

	Unk1 int32

	Addr5 uint32 `json:"-"`
	Addr6 uint32 `json:"-"`

	UnkByte1 int8

	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`

	Unk2 int8
	Unk3 int8
	Unk4 int8

	Addr11 uint32 `json:"-"`
	Addr12 uint32 `json:"-"`

	UnkByte2 int8

	Hash1 Hash
	Hash2 Hash
	Hash3 Hash
	Str1  StrWithPrefix16

	Hash4 Hash
	Hash5 Hash

	Str2  StrWithPrefix16
	Hash6 Hash
	Hash7 Hash
	Str3  StrWithPrefix16

	Str4 StrWithPrefix16
	Str5 StrWithPrefix16
}
