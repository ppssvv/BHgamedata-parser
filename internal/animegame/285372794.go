package animegame

import "encoding/json"

// 285372794

type Data_285372794 struct {
	Filesize   uint32    `json:"-"`
	EntryCount uint32    `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []WeirdID `json:"-"`
	Junk2      []uint32  `json:"-"`
	Data       []Data_285372794_Entry
}

func (s *Data_285372794) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_285372794_Entry struct {
	ID    WeirdID
	Unk1  [3]int16
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Unk2  int32
	Addr5 uint32 `json:"-"`
	Unk3  WeirdID
	Addr6 uint32 `json:"-"`
	Unk4  int16

	Hash1   Hash
	UnkStr1 StrWithPrefix16
	Hash2   Hash
	Hash3   Hash
	Hash4   Hash
	UnkStr2 StrWithPrefix16
}

type WeirdID struct {
	Part1 int32
	Part2 int32
	Part3 int32
}
