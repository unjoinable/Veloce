package coordinate

import "Veloce/internal/objects/block"

// Point represents a 3D point interface
type Point interface {
	// X returns the X coordinate
	X() float64

	// Y returns the Y coordinate
	Y() float64

	// Z returns the Z coordinate
	Z() float64

	// BlockX returns the floored X coordinate as block position
	BlockX() int

	// BlockY returns the floored Y coordinate as block position
	BlockY() int

	// BlockZ returns the floored Z coordinate as block position
	BlockZ() int

	// ChunkX returns the chunk X coordinate containing this point
	ChunkX() int

	// Section returns the chunk Y coordinate (section) containing this point
	Section() int

	// ChunkZ returns the chunk Z coordinate containing this point
	ChunkZ() int

	// WithX creates a new point with the specified X coordinate
	WithX(x float64) Point

	// WithY creates a new point with the specified Y coordinate
	WithY(y float64) Point

	// WithZ creates a new point with the specified Z coordinate
	WithZ(z float64) Point

	// Add returns a new point with the specified coordinates added
	Add(x, y, z float64) Point

	// AddPoint returns a new point with another point's coordinates added
	AddPoint(point Point) Point

	// AddValue returns a new point with the same value added to all coordinates
	AddValue(value float64) Point

	// Sub returns a new point with the specified coordinates subtracted
	Sub(x, y, z float64) Point

	// SubPoint returns a new point with another point's coordinates subtracted
	SubPoint(point Point) Point

	// SubValue returns a new point with the same value subtracted from all coordinates
	SubValue(value float64) Point

	// Mul returns a new point with coordinates multiplied by the specified values
	Mul(x, y, z float64) Point

	// MulPoint returns a new point with coordinates multiplied by another point's coordinates
	MulPoint(point Point) Point

	// MulValue returns a new point with all coordinates multiplied by the same value
	MulValue(value float64) Point

	// Div returns a new point with coordinates divided by the specified values
	Div(x, y, z float64) Point

	// DivPoint returns a new point with coordinates divided by another point's coordinates
	DivPoint(point Point) Point

	// DivValue returns a new point with all coordinates divided by the same value
	DivValue(value float64) Point

	// Relative returns a new point offset by one block in the specified face direction
	Relative(face block.Face) Point

	// DistanceSquared returns the squared distance to the specified coordinates
	DistanceSquared(x, y, z float64) float64

	// DistanceSquaredToPoint returns the squared distance to another point
	DistanceSquaredToPoint(point Point) float64

	// Distance returns the Euclidean distance to the specified coordinates
	Distance(x, y, z float64) float64

	// DistanceToPoint returns the Euclidean distance to another point
	DistanceToPoint(point Point) float64

	// SamePoint checks if this point has the same coordinates as the specified values
	SamePoint(x, y, z float64) bool

	// SamePointAs checks if this point has the same coordinates as another point
	SamePointAs(point Point) bool

	// IsZero checks if all coordinates are zero
	IsZero() bool

	// SameChunk checks if this point is in the same chunk as another point
	SameChunk(point Point) bool

	// SameBlock checks if this point is in the same block as the specified block coordinates
	SameBlock(blockX, blockY, blockZ int) bool

	// SameBlockAs checks if this point is in the same block as another point
	SameBlockAs(point Point) bool
}
