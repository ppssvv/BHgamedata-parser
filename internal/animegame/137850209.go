package animegame

import "encoding/json"

// looks like something about apho

type Data_137850209 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_137850209_Entry
}

func (s *Data_137850209) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_137850209_Entry struct {
	ID int32

	Unk1     int32
	UnkByte1 int8
	Unk2     int32
	Unk3     int32
	Unk4     int32
	Unk5     int32

	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	UnkByte2 int8

	Addr6 uint32 `json:"-"`

	UnkStr1  StrWithPrefix16
	UnkArr1  ArrStrWithPrefix16
	UnkHash1 Hash
	UnkStr2  StrWithPrefix16
	UnkArr2  ArrInt32

	UnkArr3 ArrInt32
}
