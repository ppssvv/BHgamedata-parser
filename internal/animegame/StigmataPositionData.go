package animegame

import "encoding/json"

type StigPosData struct {
	Filesize   uint32            `json:"-"`
	EntryCount uint32            `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []StrWithPrefix16 `json:"-"`
	Junk2      []uint32          `json:"-"`
	Data       []StigPosDataEntry
}

func (s *StigPosData) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type StigPosDataEntry struct {
	Addr1 uint32 `json:"-"`

	RootHeight float32
	RootWidth  float32
	HasBack    int8
	BackPosX   float32
	BackPosY   float32
	BackHeight float32
	BackWidth  float32

	Addr2 uint32 `json:"-"`

	ImageHeight float32
	ImageWidth  float32

	Addr3 uint32 `json:"-"`

	HasFront       int8
	FrontPosX      float32
	FrontPosY      float32
	FrontHeight    float32
	FrontWidth     float32
	FrontRotationX float32
	FrontRotationY float32
	FrontRotationZ float32

	Addr4 uint32 `json:"-"`

	Name      StrWithPrefix16
	BackPath  StrWithPrefix16
	ImagePath StrWithPrefix16
	FrontPath StrWithPrefix16
}
