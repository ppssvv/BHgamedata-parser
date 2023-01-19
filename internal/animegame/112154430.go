package animegame

import (
	"dataparse/internal/binreader"
	"encoding/binary"
)

type Data_112154430 []Data_112154430_Entry

type Data_112154430_Entry struct {
	ID       int64
	UnkArr1  []int32
	Unk1     int32
	UnkStr1  string
	UnkByte1 int8
}

func NewData_112154430(f string) Data_112154430 {
	reader := binreader.NewUnpacker(binary.LittleEndian, mustOpenFile(f))

	reader.Skip(4) // filesize
	entryCount := reader.Int32()
	reader.Skip(int64(entryCount) * 8) // list of hashes
	reader.Skip(int64(entryCount) * 4) // list of addresses

	result := []Data_112154430_Entry{}

	for i := 0; i < int(entryCount); i++ {
		result = append(result, newData_112154430_Entry(reader))
	}

	return result
}

func newData_112154430_Entry(reader *binreader.Unpacker) Data_112154430_Entry {
	result := Data_112154430_Entry{}

	result.ID = reader.Int64() // probably just 2 int32

	reader.Skip(4) // addr

	result.Unk1 = reader.Int32()

	reader.Skip(4) // addr

	result.UnkByte1 = reader.Int8()

	result.UnkArr1 = reader.ArrayInt32(uint64(reader.Int32()))
	result.UnkStr1 = reader.StringWithUint16Prefix()

	return result
}
