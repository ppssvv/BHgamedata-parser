package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
)

type OWActivityBossData []OWActivityBossDataEntry

type OWActivityBossDataEntry struct {
	BossId int32

	MonsterIdList []int32

	HardLevel int32
	MonsterHp int32

	ClueEventIdList   []int32
	ClueLocationId    []int32
	ClueTimeLimitTips Hash

	StaminaCost    int32
	ClueProgress   int32
	ClueTimeLimit  int32
	ClueBuffId     int32
	MutationBuffId int32
	ActivityExp    int32
	BossReward     int32

	BossName         Hash
	BossDescription  Hash
	BossLocationList []int32
	ImagePath        string

	StageId        int32
	MPDamageRatio  float32
	StageMonsterId int32
}

func NewOWActivityBossData(f string) OWActivityBossData {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 4) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []OWActivityBossDataEntry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newOWActivityBossDataEntry(reader))
	}

	return result
}

func newOWActivityBossDataEntry(reader *binreader.Unpacker) OWActivityBossDataEntry {
	result := OWActivityBossDataEntry{}

	result.BossId = reader.Int32()

	reader.Skip(4)

	result.HardLevel = reader.Int32()
	result.MonsterHp = reader.Int32()

	reader.Skip(4)
	reader.Skip(4)

	reader.Skip(4)

	result.StaminaCost = reader.Int32()
	result.ClueProgress = reader.Int32()
	result.ClueTimeLimit = reader.Int32()
	result.ClueBuffId = reader.Int32()
	result.MutationBuffId = reader.Int32()
	result.ActivityExp = reader.Int32()
	result.BossReward = reader.Int32()

	reader.Skip(4)
	reader.Skip(4)
	reader.Skip(4)
	reader.Skip(4)

	result.StageId = reader.Int32()
	result.MPDamageRatio = reader.Float32()
	result.StageMonsterId = reader.Int32()

	result.MonsterIdList = reader.ArrayInt32(uint64(reader.Uint32()))
	result.ClueEventIdList = reader.ArrayInt32(uint64(reader.Uint32()))
	result.ClueLocationId = reader.ArrayInt32(uint64(reader.Uint32()))

	result.ClueTimeLimitTips = readHash(reader)
	result.BossName = readHash(reader)
	result.BossDescription = readHash(reader)

	result.BossLocationList = reader.ArrayInt32(uint64(reader.Uint32()))

	result.ImagePath = reader.StringWithUint16Prefix()

	return result
}
