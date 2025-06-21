package codec

// Encoder interface for types that can be encoded to a transcoder format
type Encoder[T any] interface {
	Encode(coder Transcoder[any], value T) (any, bool)
}
