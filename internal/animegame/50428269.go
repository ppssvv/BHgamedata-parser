package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
)

// Looks like HybridSiteData, but not sure

type Data_50428269 []Data_50428269_Entry

type Data_50428269_Entry struct {
	SiteID int32

	ChapterID int32
	ActID     int32
	SiteType  int32
	PreSite   int32

	Arr1 []Data_50428269_

	Unk5 int32

	SiteContentID []int32
	Hash1         Hash
	Str1          string
	Str2          string

	RewardDisplayType int8
	Unk6              int32

	DisplayTimeLimit string
}

type Data_50428269_ struct {
	FieldA int32
	FieldB int32
}

func NewData_50428269(f string) Data_50428269 {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 4) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []Data_50428269_Entry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newData_50428269_Entry(reader))
	}

	return result
}

func newData_50428269_Entry(reader *binreader.Unpacker) Data_50428269_Entry {
	result := Data_50428269_Entry{}

	result.SiteID = reader.Int32()

	result.ChapterID = reader.Int32()
	result.ActID = reader.Int32()
	result.SiteType = reader.Int32()
	result.PreSite = reader.Int32()

	reader.Skip(4)

	result.Unk5 = reader.Int32()

	reader.Skip(4)
	reader.Skip(4)
	reader.Skip(4)
	reader.Skip(4)

	result.RewardDisplayType = reader.Int8()
	result.Unk6 = reader.Int32()

	reader.Skip(4)

	count := reader.Int32()
	arr := []Data_50428269_{}
	for i := int32(0); i < count; i++ {
		tmp := Data_50428269_{}

		tmp.FieldA = reader.Int32()
		tmp.FieldB = reader.Int32()

		arr = append(arr, tmp)
	}
	result.Arr1 = arr

	result.SiteContentID = reader.ArrayInt32(uint64(reader.Int32()))

	result.Hash1 = readHash(reader)

	result.Str1 = reader.StringWithUint16Prefix()
	result.Str2 = reader.StringWithUint16Prefix()
	result.DisplayTimeLimit = reader.StringWithUint16Prefix()

	return result
}
