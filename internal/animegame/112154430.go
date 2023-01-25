package animegame

import "encoding/json"

type Data_112154430 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint64 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_112154430_Entry
}

func (s *Data_112154430) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_112154430_Entry struct {
	ID       int64
	Addr1    uint32 `json:"-"`
	Unk1     int32
	Addr2    uint32 `json:"-"`
	UnkByte1 int8

	UnkArr1 ArrInt32
	UnkStr1 StrWithPrefix16
}
