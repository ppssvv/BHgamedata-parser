package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
	"encoding/json"
)

type DormFurniture []DormFurnitureEntry

// ProcessDormFurniture provides a unified interface for batch processing
func ProcessDormFurniture(f string) ([]byte, error) {
	return json.MarshalIndent(NewDormFurniture(f), "", "  ")
}

type DormFurnitureEntry struct {
	Hash                    int32
	Rarity                  int8
	Tags                    []int8
	Type                    int8
	EditType                int8
	Unk4                    int8
	InitialUnlock           int8
	SuitID                  int8
	Size                    []int8
	ManualScale             float32
	ComfortType             int8
	ComfortValue            int8
	HcoinCost               int32
	CostItemId              int32
	CostItemNum             int32
	UnlockItemId            int32
	UnlockItemNum           int32
	Unk7                    int8
	FurnitureModPath        string
	FurnitureNameText       int32
	FurnitureIconPath       string
	FurnitureDescText       int32
	FurnitureSourceDescText int32
	Unk8                    int8
	DeleteMaterialList      []material
	VerandaPath             string
	Unk9                    int8
}

func NewDormFurniture(f string) DormFurniture {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))
	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 4) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []DormFurnitureEntry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newDormFurnitureEntry(reader))
	}

	return result
}

func newDormFurnitureEntry(reader *binreader.Unpacker) DormFurnitureEntry {
	result := DormFurnitureEntry{}

	// - Header

	result.Hash = reader.Int32()
	result.Rarity = reader.Int8()

	reader.Skip(4) // addrto tags

	result.Type = reader.Int8()
	result.EditType = reader.Int8()
	result.Unk4 = reader.Int8()
	result.InitialUnlock = reader.Int8()
	result.SuitID = reader.Int8()

	reader.Skip(4) // addrto size

	result.ManualScale = reader.Float32()

	result.ComfortType = reader.Int8()
	result.ComfortValue = reader.Int8()

	result.HcoinCost = reader.Int32()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.Unk7 = reader.Int8()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.Unk8 = reader.Int8()

	reader.Skip(4) // addrto
	reader.Skip(4) // addrto

	result.Unk9 = reader.Int8()

	// - Body

	result.Tags = reader.ArrayInt8(uint64(reader.Uint32()))
	result.Size = reader.ArrayInt8(uint64(reader.Uint32()))

	hasCost := reader.Uint32()
	if hasCost > 0 {
		if hasCost > 1 {
			panic("hi")
		}

		result.CostItemId = reader.Int32()  // CostItemId
		result.CostItemNum = reader.Int32() // CostItemNum
	}

	hasUnlock := reader.Uint32()
	if hasUnlock > 0 {
		if hasUnlock > 1 {
			panic("hi2")
		}

		result.UnlockItemId = reader.Int32()  // UnlockItemId
		result.UnlockItemNum = reader.Int32() // UnlockItemNum
	}

	result.FurnitureModPath = reader.StringWithUint16Prefix()
	result.FurnitureNameText = readHash(reader)
	result.FurnitureIconPath = reader.StringWithUint16Prefix()
	result.FurnitureDescText = readHash(reader)
	result.FurnitureSourceDescText = readHash(reader)

	materialCount := reader.Int32()
	deleteMaterialList := []material{}
	for i := 0; i < int(materialCount); i++ {
		tmp := material{}
		tmp.MaterialID = reader.Int32()
		tmp.MaterialNumber = reader.Int32()

		deleteMaterialList = append(deleteMaterialList, tmp)
	}
	result.DeleteMaterialList = deleteMaterialList

	result.VerandaPath = reader.StringWithUint16Prefix()

	return result
}

type material struct {
	MaterialID     int32
	MaterialNumber int32
}
