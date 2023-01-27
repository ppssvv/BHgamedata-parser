package animegame

import "encoding/json"

// 354247105

type Data_354247105 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_354247105_Entry
}

func (s *Data_354247105) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_354247105_Entry struct {
	ID    Hash
	Count int32 `json:"-" bin:"sizeof=Arr1"`
	Arr1  []Data_354247105_Val
}

type Data_354247105_Val struct {
	ValueA float32
	ValueB float32
}
