package binreader

import (
	"encoding/binary"
	"io"
	"math"
)

// Unpacker helps you unpack binary data from an io.Reader.
type Unpacker struct {
	reader io.ReadSeeker
	endian binary.ByteOrder
}

// NewUnpacker returns a *Unpacker which hold an io.Reader. User must provide the byte order explicitly.
func NewUnpacker(endian binary.ByteOrder, reader io.ReadSeeker) *Unpacker {
	return &Unpacker{
		reader: reader,
		endian: endian,
	}
}

func (u *Unpacker) SetEndian(order binary.ByteOrder) {
	u.endian = order
}

func (u *Unpacker) Skip(n int64) {
	_, err := u.reader.Seek(n, 1)
	if err != nil {
		panic(err)
	}
}

// Byte fetch the first byte in io.Reader. Returns a byte and an error if
// exists.
func (u *Unpacker) Byte() byte {
	buffer := make([]byte, 1)
	mustRead(u.reader, buffer)
	return buffer[0]
}

// Bytes fetch n bytes in io.Reader. Returns a byte array and an error if
// exists.
func (u *Unpacker) Bytes(_n uint64) []byte {
	buffer := make([]byte, _n)
	mustRead(u.reader, buffer)
	return buffer
}

// Uint8 fetch 1 byte in io.Reader and convert it to uint8
func (u *Unpacker) Uint8() uint8 {
	buffer := make([]byte, 1)
	mustRead(u.reader, buffer)
	return buffer[0]
}

// ShiftInt16 fetch 2 bytes in io.Reader and convert it to int16.
func (u *Unpacker) Int8() int8 {
	return int8(u.Uint8())
}

// Uint16 fetch 2 bytes in io.Reader and convert it to uint16.
func (u *Unpacker) Uint16() uint16 {
	buffer := make([]byte, 2)
	mustRead(u.reader, buffer)
	return u.endian.Uint16(buffer)
}

// Int16 fetch 2 bytes in io.Reader and convert it to int16.
func (u *Unpacker) Int16() int16 {
	return int16(u.Uint16())
}

// Uint32 fetch 4 bytes in io.Reader and convert it to uint32.
func (u *Unpacker) Uint32() uint32 {
	buffer := make([]byte, 4)
	mustRead(u.reader, buffer)
	return u.endian.Uint32(buffer)
}

// Int32 fetch 4 bytes in io.Reader and convert it to int32.
func (u *Unpacker) Int32() int32 {
	return int32(u.Uint32())
}

// Uint64 fetch 8 bytes in io.Reader and convert it to uint64.
func (u *Unpacker) Uint64() uint64 {
	buffer := make([]byte, 8)
	mustRead(u.reader, buffer)
	return u.endian.Uint64(buffer)
}

// Int64 fetch 8 bytes in io.Reader and convert it to int64.
func (u *Unpacker) Int64() int64 {
	return int64(u.Uint64())
}

// Float32 fetch 4 bytes in io.Reader and convert it to float32
func (u *Unpacker) Float32() float32 {
	buffer := make([]byte, 4)
	mustRead(u.reader, buffer)
	return math.Float32frombits(u.endian.Uint32(buffer))
}

// Float64 fetch 8 bytes in io.Reader and convert it to float64
func (u *Unpacker) Float64() float64 {
	buffer := make([]byte, 8)
	mustRead(u.reader, buffer)
	return math.Float64frombits(u.endian.Uint64(buffer))
}

// ShiftString fetch n bytes, convert it to string
func (u *Unpacker) String(n uint64) string {
	buffer := make([]byte, n)
	mustRead(u.reader, buffer)
	return string(buffer)
}

// StringWithUint16Prefix read 2 bytes as string length, then read N bytes,
// convert it to string
func (u *Unpacker) StringWithUint16Prefix() string {
	return u.String(uint64(u.Uint16()))
}

// StringWithUint32Prefix read 4 bytes as string length, then read N bytes,
// convert it to string
func (u *Unpacker) StringWithUint32Prefix() string {
	return u.String(uint64(u.Uint32()))
}

// StringWithUint64Prefix read 8 bytes as string length, then read N bytes,
// convert it to string
func (u *Unpacker) StringWithUint64Prefix() string {
	return u.String(uint64(u.Uint64()))
}

func mustRead(reader io.Reader, buffer []byte) {
	if _, err := io.ReadFull(reader, buffer); err != nil {
		panic(err)
	}
}
