package codec

// RawValue interface for raw codec values
type RawValue struct {
	convert func(to Transcoder[any]) (any, error)
}

type Codec[T any] struct{}
