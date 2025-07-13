package interfaces

import (
	"Veloce/internal/objects/protocol"
	"encoding/binary"
	"errors"
	"github.com/google/uuid"
	"math"
)

// Common errors returned by Buffer operations
var (
	ErrBufferUnderflow = errors.New("buffer underflow: not enough data to read")
	ErrValueTooLarge   = errors.New("value too large: exceeds maximum allowed size")
	ErrNegativeLength  = errors.New("negative length: string length cannot be negative")
	ErrInvalidPosition = errors.New("invalid position: position out of bounds")
)

const (
	defaultCapacity = 64
	maxVarIntBytes  = 5
	maxVarLongBytes = 10
	uuidByteLength  = 16
)

// Buffer provides efficient binary data reading and writing with automatic growth.
type Buffer struct {
	data []byte
	pos  int
}

// NewBuffer creates a new Buffer instance.
func NewBuffer(data []byte) *Buffer {
	if data == nil {
		data = make([]byte, 0, defaultCapacity)
	}
	return &Buffer{data: data}
}

// Data returns the underlying byte slice containing all written data.
func (b *Buffer) Data() []byte {
	return b.data
}

// Len returns the total length of data in the buffer.
func (b *Buffer) Len() int {
	return len(b.data)
}

// Remaining returns the number of bytes available for reading from current position.
func (b *Buffer) Remaining() int {
	return len(b.data) - b.pos
}

// Position returns the current read/write position in the buffer.
func (b *Buffer) Position() int {
	return b.pos
}

// SetPosition sets the current position for reading/writing operations.
func (b *Buffer) SetPosition(pos int) error {
	if pos < 0 || pos > len(b.data) {
		return ErrInvalidPosition
	}
	b.pos = pos
	return nil
}

// Reset resets the position to the beginning of the buffer.
func (b *Buffer) Reset() {
	b.pos = 0
}

// grow ensures the buffer has enough capacity for n additional bytes.
func (b *Buffer) grow(n int) {
	required := b.pos + n
	if required <= len(b.data) {
		return
	}

	capacity := cap(b.data)
	for capacity < required {
		if capacity == 0 {
			capacity = defaultCapacity
		} else {
			capacity *= 2
		}
	}

	newData := make([]byte, required, capacity)
	copy(newData, b.data[:b.pos])
	b.data = newData
}

// checkRead verifies that n bytes are available for reading.
func (b *Buffer) checkRead(n int) error {
	if b.pos+n > len(b.data) {
		return ErrBufferUnderflow
	}
	return nil
}

// WriteByte writes a single byte to the buffer.
func (b *Buffer) WriteByte(v byte) error {
	b.grow(1)
	b.data[b.pos] = v
	b.pos++
	return nil
}

// ReadByte reads a single byte from the buffer.
func (b *Buffer) ReadByte() (byte, error) {
	if err := b.checkRead(1); err != nil {
		return 0, err
	}
	v := b.data[b.pos]
	b.pos++
	return v, nil
}

// WriteBool writes a boolean value as a single byte (1 for true, 0 for false).
func (b *Buffer) WriteBool(v bool) {
	if v {
		b.WriteByte(1)
	} else {
		b.WriteByte(0)
	}
}

// ReadBool reads a boolean value from the buffer.
func (b *Buffer) ReadBool() (bool, error) {
	v, err := b.ReadByte()
	return v != 0, err
}

// WriteInt16 writes a 16-bit signed integer in big-endian format.
func (b *Buffer) WriteInt16(v int16) {
	b.grow(2)
	binary.BigEndian.PutUint16(b.data[b.pos:], uint16(v))
	b.pos += 2
}

// ReadInt16 reads a 16-bit signed integer in big-endian format.
func (b *Buffer) ReadInt16() (int16, error) {
	if err := b.checkRead(2); err != nil {
		return 0, err
	}
	v := int16(binary.BigEndian.Uint16(b.data[b.pos:]))
	b.pos += 2
	return v, nil
}

// WriteUint16 writes a 16-bit unsigned integer in big-endian format.
func (b *Buffer) WriteUint16(v uint16) {
	b.grow(2)
	binary.BigEndian.PutUint16(b.data[b.pos:], v)
	b.pos += 2
}

// ReadUint16 reads a 16-bit unsigned integer in big-endian format.
func (b *Buffer) ReadUint16() (uint16, error) {
	if err := b.checkRead(2); err != nil {
		return 0, err
	}
	v := binary.BigEndian.Uint16(b.data[b.pos:])
	b.pos += 2
	return v, nil
}

// WriteInt32 writes a 32-bit signed integer in big-endian format.
func (b *Buffer) WriteInt32(v int32) {
	b.grow(4)
	binary.BigEndian.PutUint32(b.data[b.pos:], uint32(v))
	b.pos += 4
}

// ReadInt32 reads a 32-bit signed integer in big-endian format.
func (b *Buffer) ReadInt32() (int32, error) {
	if err := b.checkRead(4); err != nil {
		return 0, err
	}
	v := int32(binary.BigEndian.Uint32(b.data[b.pos:]))
	b.pos += 4
	return v, nil
}

// WriteUint32 writes a 32-bit unsigned integer in big-endian format.
func (b *Buffer) WriteUint32(v uint32) {
	b.grow(4)
	binary.BigEndian.PutUint32(b.data[b.pos:], v)
	b.pos += 4
}

// ReadUint32 reads a 32-bit unsigned integer in big-endian format.
func (b *Buffer) ReadUint32() (uint32, error) {
	if err := b.checkRead(4); err != nil {
		return 0, err
	}
	v := binary.BigEndian.Uint32(b.data[b.pos:])
	b.pos += 4
	return v, nil
}

// WriteInt64 writes a 64-bit signed integer in big-endian format.
func (b *Buffer) WriteInt64(v int64) {
	b.grow(8)
	binary.BigEndian.PutUint64(b.data[b.pos:], uint64(v))
	b.pos += 8
}

// ReadInt64 reads a 64-bit signed integer in big-endian format.
func (b *Buffer) ReadInt64() (int64, error) {
	if err := b.checkRead(8); err != nil {
		return 0, err
	}
	v := int64(binary.BigEndian.Uint64(b.data[b.pos:]))
	b.pos += 8
	return v, nil
}

// WriteUint64 writes a 64-bit unsigned integer in big-endian format.
func (b *Buffer) WriteUint64(v uint64) {
	b.grow(8)
	binary.BigEndian.PutUint64(b.data[b.pos:], v)
	b.pos += 8
}

// ReadUint64 reads a 64-bit unsigned integer in big-endian format.
func (b *Buffer) ReadUint64() (uint64, error) {
	if err := b.checkRead(8); err != nil {
		return 0, err
	}
	v := binary.BigEndian.Uint64(b.data[b.pos:])
	b.pos += 8
	return v, nil
}

// WriteFloat32 writes a 32-bit IEEE 754 floating-point number.
func (b *Buffer) WriteFloat32(v float32) {
	b.WriteUint32(math.Float32bits(v))
}

// ReadFloat32 reads a 32-bit IEEE 754 floating-point number.
func (b *Buffer) ReadFloat32() (float32, error) {
	bits, err := b.ReadUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(bits), nil
}

// WriteFloat64 writes a 64-bit IEEE 754 floating-point number.
func (b *Buffer) WriteFloat64(v float64) {
	b.WriteUint64(math.Float64bits(v))
}

// ReadFloat64 reads a 64-bit IEEE 754 floating-point number.
func (b *Buffer) ReadFloat64() (float64, error) {
	bits, err := b.ReadUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(bits), nil
}

// WriteVarInt writes a Minecraft VarInt (variable-length encoded int32) to the buffer.
func (b *Buffer) WriteVarInt(v protocol.VarInt) error {
	uv := uint32(v)

	for {
		temp := byte(uv & 0x7F)
		uv >>= 7

		if uv != 0 {
			temp |= 0x80
		}

		b.WriteByte(temp)

		if uv == 0 {
			break
		}
	}

	return nil
}

// ReadVarInt reads a Minecraft VarInt (variable-length encoded int32) from the buffer.
// It returns the value as a protocol.VarInt.
func (b *Buffer) ReadVarInt() (protocol.VarInt, error) {
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

	return protocol.VarInt(int32(result)), nil
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
		b.WriteByte(temp)
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

// WriteBytes writes a byte slice to the buffer.
func (b *Buffer) WriteBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	b.grow(len(data))
	copy(b.data[b.pos:], data)
	b.pos += len(data)
}

// ReadBytes reads exactly n bytes from the buffer.
func (b *Buffer) ReadBytes(n int) ([]byte, error) {
	if err := b.checkRead(n); err != nil {
		return nil, err
	}
	data := make([]byte, n)
	copy(data, b.data[b.pos:b.pos+n])
	b.pos += n
	return data, nil
}

// WriteString writes a UTF-8 string with variable-length prefix encoding.
func (b *Buffer) WriteString(s string) error {
	data := []byte(s)
	length := protocol.VarInt(len(data))

	if err := b.WriteVarInt(length); err != nil {
		return err
	}

	b.WriteBytes(data)
	return nil
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
	data, err := b.ReadBytes(int(length))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteUUID writes a UUID as 16 bytes in its canonical binary representation.
func (b *Buffer) WriteUUID(u uuid.UUID) {
	b.WriteBytes(u[:])
}

// ReadUUID reads a UUID from exactly 16 bytes in binary representation.
func (b *Buffer) ReadUUID() (uuid.UUID, error) {
	data, err := b.ReadBytes(uuidByteLength)
	if err != nil {
		return uuid.UUID{}, err
	}

	var u uuid.UUID
	copy(u[:], data)
	return u, nil
}
