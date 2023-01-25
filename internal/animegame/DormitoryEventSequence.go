package animegame

import (
	"encoding/json"
)

type DormEvent struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []DormEventEntry
}

func (s *DormEvent) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type DormEventEntry struct {
	Hash          int32
	Addr1         uint32 `json:"-"`
	Avatar        int8
	Wait          float32
	WaitRandomAdd float32
	Addr2         uint32 `json:"-"`
	TalkToAvatar  int8
	Unknown       int8
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Addr5         uint32 `json:"-"`

	Type             StrWithPrefix16
	Emotion          StrWithPrefix16
	TalkTxt          StrWithPrefix16
	FaceAnimType     StrWithPrefix16
	TriggerAction    StrWithPrefix16
	TriggerSubAction StrWithPrefix16
}
