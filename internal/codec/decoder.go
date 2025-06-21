package codec

// Decoder interface for types that can be decoded from a transcoder format
type Decoder[T any] interface {
	Decode(coder Transcoder[any], value any) (T, bool)
}
