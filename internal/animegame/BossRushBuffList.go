package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
)

type BossRushBuffList []BossRushBuffListEntry

type BossRushBuffListEntry struct {
	ID          int32
	Description Hash
	Name        Hash
	AbilityName string
	IconPath    string
	Unk1        int32
}

func NewBossRushBuffList(f string) BossRushBuffList {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 4) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []BossRushBuffListEntry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newBossRushBuffListEntry(reader))
	}

	return result
}

func newBossRushBuffListEntry(reader *binreader.Unpacker) BossRushBuffListEntry {
	result := BossRushBuffListEntry{}

	result.ID = reader.Int32()

	reader.Skip(4 * 4)

	result.Unk1 = reader.Int32()

	result.Name = readHash(reader)
	result.Description = readHash(reader)
	result.AbilityName = reader.StringWithUint16Prefix()
	result.IconPath = reader.StringWithUint16Prefix()

	return result
}
