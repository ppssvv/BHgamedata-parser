package animegame

import (
	"encoding/binary"
	"encoding/json"

	bin "github.com/streamingfast/binary"
)

// Looks like HybridSiteData, but not sure

type Data_50428269 struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []Data_50428269_Entry
}

func (s *Data_50428269) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type Data_50428269_Entry struct {
	SiteID int32

	ChapterID int32
	ActID     int32
	SiteType  int32
	PreSite   int32

	Addr1 uint32 `json:"-"`

	Unk5 int32

	Addr2 uint32 `json:"-"`
	Addr3 uint32 `json:"-"`
	Addr4 uint32 `json:"-"`
	Addr5 uint32 `json:"-"`

	RewardDisplayType int8
	Unk6              int32

	Addr6 uint32 `json:"-"`

	Arr1 Data_50428269_

	SiteContentID ArrInt32
	Hash1         Hash
	Str1          StrWithPrefix16
	Str2          StrWithPrefix16

	DisplayTimeLimit StrWithPrefix16
}

type Data_50428269_ []struct {
	FieldA int32
	FieldB int32
}

func (d *Data_50428269_) UnmarshalBinary(decoder *bin.Decoder) error {
	count, err := decoder.ReadInt32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*d = make([]struct {
		FieldA int32
		FieldB int32
	}, count)

	for i := int32(0); i < count; i++ {
		a, err := decoder.ReadInt32(binary.LittleEndian)
		if err != nil {
			return err
		}
		b, err := decoder.ReadInt32(binary.LittleEndian)
		if err != nil {
			return err
		}

		(*d)[i] = struct {
			FieldA int32
			FieldB int32
		}{a, b}
	}

	return nil
}
