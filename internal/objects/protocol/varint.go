// Package protocol defines Minecraft-specific protocol types and helpers.
package protocol

// VarInt represents a Minecraft-style variable-length integer (VarInt).
//
// VarInt is encoded using 1 to 5 bytes depending on the value, where each byte
// uses the most significant bit (MSB) as a continuation flag. This format is
// used extensively in the Minecraft protocol for compact integer encoding.
//
// For example:
//   - 0           => [0x00]
//   - 127         => [0x7F]
//   - 128         => [0x80 0x01]
//   - 25565       => [0xDD 0xC7 0x01]
//   - -1          => [0xFF 0xFF 0xFF 0xFF 0x0F]
type VarInt int32

// Int returns the raw int32 value of the VarInt.
func (v VarInt) Int() int32 {
	return int32(v)
}

// Size returns the number of bytes required to encode the VarInt.
func (v VarInt) Size() int {
	val := uint32(v)
	size := 0
	for {
		size++
		if val&^0x7F == 0 {
			break
		}
		val >>= 7
	}
	return size
}
