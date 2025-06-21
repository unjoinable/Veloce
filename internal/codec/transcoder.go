package codec

// MapLike represents a map-like structure
type MapLike[T any] interface {
	Keys() []string
	HasValue(key string) bool
	GetValue(key string) T
	Size() int
	IsEmpty() bool
}

// ListBuilder builds lists of type T
type ListBuilder[T any] interface {
	Add(value T) ListBuilder[T]
	Build() T
}

// MapBuilder builds maps of type T
type MapBuilder[T any] interface {
	Put(key T, value T) MapBuilder[T]
	PutString(key string, value T) MapBuilder[T]
	Build() T
}

// Transcoder is the main interface for transcoding between different data formats
type Transcoder[T any] interface {
	CreateNull() T

	GetBoolean(value T) (bool, bool)
	CreateBoolean(value bool) T

	GetByte(value T) (byte, bool)
	CreateByte(value byte) T

	GetInt16(value T) (int16, bool)
	CreateInt16(value int16) T

	GetInt32(value T) (int32, bool)
	CreateInt32(value int32) T

	GetInt64(value T) (int64, bool)
	CreateInt64(value int64) T

	GetFloat32(value T) (float32, bool)
	CreateFloat32(value float32) T

	GetFloat64(value T) (float64, bool)
	CreateFloat64(value float64) T

	GetString(value T) (string, bool)
	CreateString(value string) T

	GetList(value T) ([]T, bool)
	EmptyList() T
	CreateList(expectedSize int) ListBuilder[T]

	GetMap(value T) (MapLike[T], bool)
	EmptyMap() T
	CreateMap() MapBuilder[T]

	GetByteArray(value T) ([]byte, bool)
	CreateByteArray(value []byte) T

	GetInt32Array(value T) ([]int32, bool)
	CreateInt32Array(value []int32) T

	GetInt64Array(value T) ([]int64, bool)
	CreateInt64Array(value []int64) T

	ConvertTo(coder Transcoder[any], value T) (any, bool)
}
