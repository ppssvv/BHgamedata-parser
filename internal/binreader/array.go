package binreader

import "encoding/binary"

// BytesWithUint16Prefix read 2 bytes as bytes length, then read N bytes
func (u *Unpacker) BytesWithUint16Prefix() []byte {
	return u.Bytes(uint64(u.Uint16()))
}

// BytesWithUint32Prefix read 4 bytes as bytes length, then read N bytes
func (u *Unpacker) BytesWithUint32Prefix() []byte {
	return u.Bytes(uint64(u.Uint32()))
}

// BytesWithUint64Prefix read 8 bytes as bytes length, then read N bytes
func (u *Unpacker) BytesWithUint64Prefix() []byte {
	return u.Bytes(uint64(u.Uint64()))
}

func (u *Unpacker) ArrayUint8(len uint64) []uint8 {
	result := make([]uint8, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}

func (u *Unpacker) ArrayInt8(len uint64) []int8 {
	result := make([]int8, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}

func (u *Unpacker) ArrayUint16(len uint64) []uint16 {
	result := make([]uint16, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}

func (u *Unpacker) ArrayInt16(len uint64) []int16 {
	result := make([]int16, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}

func (u *Unpacker) ArrayUint32(len uint64) []uint32 {
	result := make([]uint32, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}

func (u *Unpacker) ArrayInt32(len uint64) []int32 {
	result := make([]int32, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}

func (u *Unpacker) ArrayUint64(len uint64) []uint64 {
	result := make([]uint64, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}

func (u *Unpacker) ArrayInt64(len uint64) []int64 {
	result := make([]int64, len)

	if err := binary.Read(u.reader, u.endian, &result); err != nil {
		panic(err)
	}

	return result
}
