package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
)

// looks like something about apho

type Data_137850209 []Data_137850209_Entry

type Data_137850209_Entry struct {
	ID int32

	Unk1     int32
	UnkByte1 int8
	Unk2     int32
	Unk3     int32
	Unk4     int32
	Unk5     int32

	UnkStr1  string
	UnkArr1  []string
	UnkHash1 Hash
	UnkStr2  string
	UnkArr2  []int32

	UnkByte2 int8

	UnkArr3 []int32
}

func NewData_137850209(f string) Data_137850209 {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 4) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []Data_137850209_Entry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newData_137850209_Entry(reader))
	}

	return result
}

func newData_137850209_Entry(reader *binreader.Unpacker) Data_137850209_Entry {
	result := Data_137850209_Entry{}

	result.ID = reader.Int32()

	result.Unk1 = reader.Int32()
	result.UnkByte1 = reader.Int8()

	result.Unk2 = reader.Int32()
	result.Unk3 = reader.Int32()
	result.Unk4 = reader.Int32()
	result.Unk5 = reader.Int32()

	reader.Skip(4 * 5)

	result.UnkByte2 = reader.Int8()

	reader.Skip(4)

	result.UnkStr1 = reader.StringWithUint16Prefix()

	count := reader.Int32()
	strarr := []string{}
	for i := int32(0); i < count; i++ {
		strarr = append(strarr, reader.StringWithUint16Prefix())
	}
	result.UnkArr1 = strarr

	result.UnkHash1 = readHash(reader)
	result.UnkStr2 = reader.StringWithUint16Prefix()
	result.UnkArr2 = reader.ArrayInt32(uint64(reader.Int32()))
	result.UnkArr3 = reader.ArrayInt32(uint64(reader.Int32()))

	return result
}
