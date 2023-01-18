package animetype

import (
	"dataparse/internal/binreader"
	"encoding/binary"
)

type WorldMap []WorldMapEntry

type WorldMapEntry struct {
	ID      int32
	EntryID int32

	FeatureType       []int32
	ContentIDList     []int32
	EntryTitle        Hash
	EntryImagePath    string
	EntryTitleImgPath string

	EntryCG          int32
	UnlockLevel      int32
	UnlockHintLevel  int32
	ForceShowContent int32
	InfoID           int32
	ShopType         int32

	EntryDesc            Hash
	EntryDetailImagePath string

	EntryRewardID   int32
	UnkByte1        byte
	VersionTagMaxLv int32
	VersionTagMinLv int32

	UnkByte2 byte
}

func NewWorldMap(f string) WorldMap {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()

	reader.Skip(int64(entryCount) * 4) // list of id's
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := make([]WorldMapEntry, 0, entryCount)

	for i := 0; i < int(entryCount); i++ {
		tmp := newWorldMapEntry(reader)
		// tmp.Name = names[i]
		result = append(result, tmp)
	}

	return result
}

func newWorldMapEntry(reader *binreader.Unpacker) WorldMapEntry {
	result := WorldMapEntry{}

	result.ID = reader.Int32()
	result.EntryID = reader.Int32()

	reader.Skip(4)
	reader.Skip(4)
	reader.Skip(4)
	reader.Skip(4)
	reader.Skip(4)

	result.EntryCG = reader.Int32()
	result.UnlockLevel = reader.Int32()
	result.UnlockHintLevel = reader.Int32()
	result.ForceShowContent = reader.Int32()
	result.InfoID = reader.Int32()
	result.ShopType = reader.Int32()

	reader.Skip(4)
	reader.Skip(4)

	result.EntryRewardID = reader.Int32()
	result.UnkByte1 = reader.Byte()
	result.VersionTagMaxLv = reader.Int32()
	result.VersionTagMinLv = reader.Int32()
	result.UnkByte2 = reader.Byte()

	// - Data
	result.FeatureType = reader.ArrayInt32(uint64(reader.Int32()))
	result.ContentIDList = reader.ArrayInt32(uint64(reader.Int32()))

	result.EntryTitle = readHash(reader)
	result.EntryImagePath = reader.StringWithUint16Prefix()
	result.EntryTitleImgPath = reader.StringWithUint16Prefix()
	result.EntryDesc = readHash(reader)
	result.EntryDetailImagePath = reader.StringWithUint16Prefix()

	return result
}
