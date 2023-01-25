package animegame

import (
	"encoding/binary"
	"encoding/json"

	bin "github.com/streamingfast/binary"
)

type DormFurniture struct {
	Filesize   uint32   `json:"-"`
	EntryCount uint32   `json:"-" bin:"sizeof=Junk1,Junk2,Data"`
	Junk1      []uint32 `json:"-"`
	Junk2      []uint32 `json:"-"`
	Data       []DormFurnitureEntry
}

func (s *DormFurniture) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}

type DormFurnitureEntry struct {
	Hash          int32
	Rarity        int8
	Addr1         uint32 `json:"-"`
	Type          int8
	EditType      int8
	Unk4          int8
	InitialUnlock int8
	SuitID        int8
	Addr2         uint32 `json:"-"`
	ManualScale   float32
	ComfortType   int8
	ComfortValue  int8
	HcoinCost     int32
	Addr3         uint32 `json:"-"`
	Addr4         uint32 `json:"-"`
	Unk7          int8
	Addr5         uint32 `json:"-"`
	Addr6         uint32 `json:"-"`
	Addr7         uint32 `json:"-"`
	Addr8         uint32 `json:"-"`
	Addr9         uint32 `json:"-"`
	Unk8          int8
	Addr10        uint32 `json:"-"`
	Addr11        uint32 `json:"-"`
	Unk9          int8

	Tags                    ArrInt8
	Size                    ArrInt8
	Cost                    DormFurnitureCost
	Unlock                  DormFurnitureCost
	FurnitureModPath        StrWithPrefix16
	FurnitureNameText       Hash
	FurnitureIconPath       StrWithPrefix16
	FurnitureDescText       Hash
	FurnitureSourceDescText Hash
	DeleteMaterialList      DormFurnitureMaterials
	VerandaPath             StrWithPrefix16
}

type DormFurnitureCost struct {
	Key     uint32 `json:",omitempty"`
	ItemId  int32
	ItemNum int32
}

func (c *DormFurnitureCost) UnmarshalBinary(decoder *bin.Decoder) error {
	key, err := decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return err
	}

	if key <= 0 {
		return nil
	}

	if key > 1 {
		c.Key = key
	}

	c.ItemId, err = decoder.ReadInt32(binary.LittleEndian)
	if err != nil {
		return err
	}

	c.ItemNum, err = decoder.ReadInt32(binary.LittleEndian)
	if err != nil {
		return err
	}

	return nil
}

type DormFurnitureMaterials []struct {
	MaterialID     int32
	MaterialNumber int32
}

func (m *DormFurnitureMaterials) UnmarshalBinary(decoder *bin.Decoder) error {
	count, err := decoder.ReadUint32(binary.LittleEndian)
	if err != nil {
		return err
	}

	*m = make([]struct {
		MaterialID     int32
		MaterialNumber int32
	}, count)

	for i := uint32(0); i < count; i++ {
		id, err := decoder.ReadInt32(binary.LittleEndian)
		if err != nil {
			return err
		}
		num, err := decoder.ReadInt32(binary.LittleEndian)
		if err != nil {
			return err
		}

		(*m)[i] = struct {
			MaterialID     int32
			MaterialNumber int32
		}{id, num}
	}

	return nil
}
