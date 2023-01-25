package animegame

import "encoding/json"

type TextMap struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []Hash   `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []TextMapEntry
}

func (s *TextMap) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type TextMapEntry struct {
	Addr1 uint32 `json:"-"`
	Addr2 uint32 `json:"-"`

	Hash Hash
	Text StrWithPrefix16
}
