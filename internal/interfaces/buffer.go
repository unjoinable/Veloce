package interfaces

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/google/uuid"
)

// Common errors returned by Buffer operations
var (
	ErrValueTooLarge  = errors.New("value too large: exceeds maximum allowed size")
	ErrNegativeLength = errors.New("negative length: string length cannot be negative")
)

const (
	maxVarIntBytes  = 5
	maxVarLongBytes = 10
	uuidByteLength  = 16
)

// Buffer wraps bytes.Buffer with additional protocol-specific methods
type Buffer struct {
	*bytes.Buffer
}

// NewBuffer creates a new Buffer instance.
func NewBuffer(data []byte) *Buffer {
	if data == nil {
		return &Buffer{&bytes.Buffer{}}
	}
	return &Buffer{bytes.NewBuffer(data)}
}

// WriteVarInt writes a Minecraft VarInt (variable-length encoded int32) to the buffer.
func (b *Buffer) WriteVarInt(v int32) error {
	uv := uint32(v)

	for {
		temp := byte(uv & 0x7F)
		uv >>= 7

		if uv != 0 {
			temp |= 0x80
		}

		if err := b.WriteByte(temp); err != nil {
			return err
		}

		if uv == 0 {
			break
		}
	}

	return nil
}

// ReadVarInt reads a VarInt (variable-length encoded int32) from the buffer.
func (b *Buffer) ReadVarInt() (int32, error) {
	var result uint32
	var numRead int

	for {
		read, err := b.ReadByte()
		if err != nil {
			return 0, err
		}

		value := uint32(read & 0x7F)
		result |= value << (7 * numRead)
		numRead++

		if numRead > maxVarIntBytes {
			return 0, ErrValueTooLarge
		}

		if (read & 0x80) == 0 {
			break
		}
	}

	return int32(result), nil
}

// WriteBool writes a boolean value as a single byte (1 for true, 0 for false).
func (b *Buffer) WriteBool(v bool) error {
	if v {
		return b.WriteByte(1)
	}
	return b.WriteByte(0)
}

// ReadBool reads a boolean value from the buffer.
func (b *Buffer) ReadBool() (bool, error) {
	v, err := b.ReadByte()
	return v != 0, err
}

// WriteString writes a UTF-8 string with variable-length prefix encoding.
func (b *Buffer) WriteString(s string) error {
	if err := b.WriteVarInt(int32(len(s))); err != nil {
		return err
	}
	_, err := b.Buffer.WriteString(s)
	return err
}

// ReadString reads a UTF-8 string with variable-length prefix encoding.
func (b *Buffer) ReadString() (string, error) {
	length, err := b.ReadVarInt()
	if err != nil {
		return "", err
	}
	if length < 0 {
		return "", ErrNegativeLength
	}

	data := make([]byte, length)
	_, err = b.Read(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteInt16 writes a 16-bit signed integer in big-endian format.
func (b *Buffer) WriteInt16(v int16) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadInt16 reads a 16-bit signed integer in big-endian format.
func (b *Buffer) ReadInt16() (int16, error) {
	var v int16
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteUint16 writes a 16-bit unsigned integer in big-endian format.
func (b *Buffer) WriteUint16(v uint16) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadUint16 reads a 16-bit unsigned integer in big-endian format.
func (b *Buffer) ReadUint16() (uint16, error) {
	var v uint16
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteInt32 writes a 32-bit signed integer in big-endian format.
func (b *Buffer) WriteInt32(v int32) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadInt32 reads a 32-bit signed integer in big-endian format.
func (b *Buffer) ReadInt32() (int32, error) {
	var v int32
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteUint32 writes a 32-bit unsigned integer in big-endian format.
func (b *Buffer) WriteUint32(v uint32) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadUint32 reads a 32-bit unsigned integer in big-endian format.
func (b *Buffer) ReadUint32() (uint32, error) {
	var v uint32
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteInt64 writes a 64-bit signed integer in big-endian format.
func (b *Buffer) WriteInt64(v int64) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadInt64 reads a 64-bit signed integer in big-endian format.
func (b *Buffer) ReadInt64() (int64, error) {
	var v int64
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteUint64 writes a 64-bit unsigned integer in big-endian format.
func (b *Buffer) WriteUint64(v uint64) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadUint64 reads a 64-bit unsigned integer in big-endian format.
func (b *Buffer) ReadUint64() (uint64, error) {
	var v uint64
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteFloat32 writes a 32-bit IEEE 754 floating-point number.
func (b *Buffer) WriteFloat32(v float32) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadFloat32 reads a 32-bit IEEE 754 floating-point number.
func (b *Buffer) ReadFloat32() (float32, error) {
	var v float32
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteFloat64 writes a 64-bit IEEE 754 floating-point number.
func (b *Buffer) WriteFloat64(v float64) error {
	return binary.Write(b, binary.BigEndian, v)
}

// ReadFloat64 reads a 64-bit IEEE 754 floating-point number.
func (b *Buffer) ReadFloat64() (float64, error) {
	var v float64
	err := binary.Read(b, binary.BigEndian, &v)
	return v, err
}

// WriteVarLong writes a variable-length encoded 64-bit integer.
func (b *Buffer) WriteVarLong(v int64) error {
	uv := uint64(v)
	for {
		temp := byte(uv & 0x7F)
		uv >>= 7
		if uv != 0 {
			temp |= 0x80
		}
		if err := b.WriteByte(temp); err != nil {
			return err
		}
		if uv == 0 {
			break
		}
	}
	return nil
}

// ReadVarLong reads a variable-length encoded 64-bit integer.
func (b *Buffer) ReadVarLong() (int64, error) {
	var result uint64
	var numRead int

	for {
		read, err := b.ReadByte()
		if err != nil {
			return 0, err
		}

		value := uint64(read & 0x7F)
		result |= value << (7 * numRead)
		numRead++

		if numRead > maxVarLongBytes {
			return 0, ErrValueTooLarge
		}

		if (read & 0x80) == 0 {
			break
		}
	}

	return int64(result), nil
}

// WriteUUID writes a UUID as 16 bytes in its canonical binary representation.
func (b *Buffer) WriteUUID(u uuid.UUID) error {
	_, err := b.Write(u[:])
	return err
}

// ReadUUID reads a UUID from exactly 16 bytes in binary representation.
func (b *Buffer) ReadUUID() (uuid.UUID, error) {
	data := make([]byte, uuidByteLength)
	_, err := b.Read(data)
	if err != nil {
		return uuid.UUID{}, err
	}

	var u uuid.UUID
	copy(u[:], data)
	return u, nil
}
