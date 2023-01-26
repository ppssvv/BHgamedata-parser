package animegame

import "encoding/json"

// 201380970

type Data_201380970 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_201380970_Entry
}

func (s *Data_201380970) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_201380970_Entry struct {
	ID     int32
	Unk1   int32
	Addr1  uint32 `json:"-"`
	Unk2   int32
	Unk3   int32
	Addr2  uint32 `json:"-"`
	Addr3  uint32 `json:"-"`
	Addr4  uint32 `json:"-"`
	Addr5  uint32 `json:"-"`
	Unk4   float32
	Addr6  uint32 `json:"-"`
	Addr7  uint32 `json:"-"`
	Addr8  uint32 `json:"-"`
	Unk5   int16
	Addr9  uint32 `json:"-"`
	Addr10 uint32 `json:"-"`
	Addr11 uint32 `json:"-"`

	Hash1 Hash
	Arr1  ArrInt32
	Str1  StrWithPrefix16
	Str2  StrWithPrefix16
	Str3  StrWithPrefix16
	Arr2  ArrFloat32
	Arr3  ArrFloat32
	Arr4  ArrInt32
	Arr5  ArrInt16
	Arr6  ArrInt16
	Arr7  ArrInt8
}
