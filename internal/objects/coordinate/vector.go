package coordinate

import (
	"Veloce/internal/objects/block"
	"math"
)

// Vector represents an immutable 3D vector
type Vector struct {
	x float64
	y float64
	z float64
}

// Common vector constants
var (
	Zero    = NewVector(0, 0, 0)
	One     = NewVector(1, 1, 1)
	Section = NewVector(16, 16, 16)
)

const Epsilon = 0.000001

// NewVector creates a new vector with the specified coordinates
func NewVector(x, y, z float64) *Vector {
	return &Vector{x: x, y: y, z: z}
}

// NewVectorXZ creates a new vector with X and Z coordinates, Y is set to 0
func NewVectorXZ(x, z float64) *Vector {
	return &Vector{x: x, y: 0, z: z}
}

// NewVectorValue creates a vector with all 3 coordinates sharing the same value
func NewVectorValue(value float64) *Vector {
	return &Vector{x: value, y: value, z: value}
}

// VectorFromPoint converts a Point into a Vector
func VectorFromPoint(point Point) *Vector {
	if v, ok := point.(*Vector); ok {
		return v
	}
	return NewVector(point.X(), point.Y(), point.Z())
}

// Interpolation defines interpolation functions
type Interpolation func(float64) float64

// Common interpolations
var (
	InterpolationLinear = func(a float64) float64 { return a }
	InterpolationSmooth = func(a float64) float64 { return a * a * (3 - 2*a) }
)

// Point interface implementation

func (v *Vector) X() float64 {
	return v.x
}

func (v *Vector) Y() float64 {
	return v.y
}

func (v *Vector) Z() float64 {
	return v.z
}

func (v *Vector) BlockX() int {
	return globalToBlock(v.x)
}

func (v *Vector) BlockY() int {
	return globalToBlock(v.y)
}

func (v *Vector) BlockZ() int {
	return globalToBlock(v.z)
}

func (v *Vector) ChunkX() int {
	return globalToChunk(v.x)
}

func (v *Vector) Section() int {
	return globalToChunk(v.y)
}

func (v *Vector) ChunkZ() int {
	return globalToChunk(v.z)
}

func (v *Vector) WithX(x float64) Point {
	return NewVector(x, v.y, v.z)
}

func (v *Vector) WithY(y float64) Point {
	return NewVector(v.x, y, v.z)
}

func (v *Vector) WithZ(z float64) Point {
	return NewVector(v.x, v.y, z)
}

func (v *Vector) Add(x, y, z float64) Point {
	return NewVector(v.x+x, v.y+y, v.z+z)
}

func (v *Vector) AddPoint(point Point) Point {
	return v.Add(point.X(), point.Y(), point.Z())
}

func (v *Vector) AddValue(value float64) Point {
	return v.Add(value, value, value)
}

func (v *Vector) Sub(x, y, z float64) Point {
	return NewVector(v.x-x, v.y-y, v.z-z)
}

func (v *Vector) SubPoint(point Point) Point {
	return v.Sub(point.X(), point.Y(), point.Z())
}

func (v *Vector) SubValue(value float64) Point {
	return v.Sub(value, value, value)
}

func (v *Vector) Mul(x, y, z float64) Point {
	return NewVector(v.x*x, v.y*y, v.z*z)
}

func (v *Vector) MulPoint(point Point) Point {
	return v.Mul(point.X(), point.Y(), point.Z())
}

func (v *Vector) MulValue(value float64) Point {
	return v.Mul(value, value, value)
}

func (v *Vector) Div(x, y, z float64) Point {
	return NewVector(v.x/x, v.y/y, v.z/z)
}

func (v *Vector) DivPoint(point Point) Point {
	return v.Div(point.X(), point.Y(), point.Z())
}

func (v *Vector) DivValue(value float64) Point {
	return v.Div(value, value, value)
}

func (v *Vector) Relative(face block.Face) Point {
	switch face {
	case block.Bottom:
		return v.Sub(0, 1, 0)
	case block.Top:
		return v.Add(0, 1, 0)
	case block.North:
		return v.Sub(0, 0, 1)
	case block.South:
		return v.Add(0, 0, 1)
	case block.West:
		return v.Sub(1, 0, 0)
	case block.East:
		return v.Add(1, 0, 0)
	default:
		return v
	}
}

func (v *Vector) DistanceSquared(x, y, z float64) float64 {
	dx := v.x - x
	dy := v.y - y
	dz := v.z - z
	return dx*dx + dy*dy + dz*dz
}

func (v *Vector) DistanceSquaredToPoint(point Point) float64 {
	return v.DistanceSquared(point.X(), point.Y(), point.Z())
}

func (v *Vector) Distance(x, y, z float64) float64 {
	return math.Sqrt(v.DistanceSquared(x, y, z))
}

func (v *Vector) DistanceToPoint(point Point) float64 {
	return v.Distance(point.X(), point.Y(), point.Z())
}

func (v *Vector) SamePoint(x, y, z float64) bool {
	return v.x == x && v.y == y && v.z == z
}

func (v *Vector) SamePointAs(point Point) bool {
	return v.SamePoint(point.X(), point.Y(), point.Z())
}

func (v *Vector) IsZero() bool {
	return v.x == 0 && v.y == 0 && v.z == 0
}

func (v *Vector) SameChunk(point Point) bool {
	return v.ChunkX() == point.ChunkX() && v.ChunkZ() == point.ChunkZ()
}

func (v *Vector) SameBlock(blockX, blockY, blockZ int) bool {
	return v.BlockX() == blockX && v.BlockY() == blockY && v.BlockZ() == blockZ
}

func (v *Vector) SameBlockAs(point Point) bool {
	return v.SameBlock(point.BlockX(), point.BlockY(), point.BlockZ())
}

// Vector-specific methods

// Neg returns the negated vector
func (v *Vector) Neg() *Vector {
	return NewVector(-v.x, -v.y, -v.z)
}

// Abs returns a vector with absolute values
func (v *Vector) Abs() *Vector {
	return NewVector(math.Abs(v.x), math.Abs(v.y), math.Abs(v.z))
}

// Floor returns a vector with each coordinate floored to the nearest integer
func (v *Vector) Floor() *Vector {
	return NewVector(math.Floor(v.x), math.Floor(v.y), math.Floor(v.z))
}

// Ceil returns a vector with each coordinate ceiled to the nearest integer
func (v *Vector) Ceil() *Vector {
	return NewVector(math.Ceil(v.x), math.Ceil(v.y), math.Ceil(v.z))
}

// Signum returns a vector with the sign of each coordinate (-1, 0, or 1)
func (v *Vector) Signum() *Vector {
	return NewVector(signum(v.x), signum(v.y), signum(v.z))
}

// Epsilon returns a vector with coordinates set to 0 if they're within epsilon of 0
func (v *Vector) Epsilon() *Vector {
	x, y, z := v.x, v.y, v.z
	if math.Abs(x) < Epsilon {
		x = 0
	}
	if math.Abs(y) < Epsilon {
		y = 0
	}
	if math.Abs(z) < Epsilon {
		z = 0
	}
	return NewVector(x, y, z)
}

// Min returns a vector with minimum values compared to another point
func (v *Vector) Min(point Point) *Vector {
	return NewVector(math.Min(v.x, point.X()), math.Min(v.y, point.Y()), math.Min(v.z, point.Z()))
}

// MinCoords returns a vector with minimum values compared to coordinates
func (v *Vector) MinCoords(x, y, z float64) *Vector {
	return NewVector(math.Min(v.x, x), math.Min(v.y, y), math.Min(v.z, z))
}

// MinValue returns a vector with minimum values compared to a single value
func (v *Vector) MinValue(value float64) *Vector {
	return NewVector(math.Min(v.x, value), math.Min(v.y, value), math.Min(v.z, value))
}

// Max returns a vector with maximum values compared to another point
func (v *Vector) Max(point Point) *Vector {
	return NewVector(math.Max(v.x, point.X()), math.Max(v.y, point.Y()), math.Max(v.z, point.Z()))
}

// MaxCoords returns a vector with maximum values compared to coordinates
func (v *Vector) MaxCoords(x, y, z float64) *Vector {
	return NewVector(math.Max(v.x, x), math.Max(v.y, y), math.Max(v.z, z))
}

// MaxValue returns a vector with maximum values compared to a single value
func (v *Vector) MaxValue(value float64) *Vector {
	return NewVector(math.Max(v.x, value), math.Max(v.y, value), math.Max(v.z, value))
}

// LengthSquared returns the magnitude of the vector squared
func (v *Vector) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// Length returns the magnitude of the vector
func (v *Vector) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// Normalize converts this vector to a unit vector (a vector with length of 1)
func (v *Vector) Normalize() *Vector {
	length := v.Length()
	if length == 0 {
		return NewVector(0, 0, 0)
	}
	return NewVector(v.x/length, v.y/length, v.z/length)
}

// IsNormalized returns if a vector is normalized
func (v *Vector) IsNormalized() bool {
	return math.Abs(v.LengthSquared()-1) < Epsilon
}

// Angle gets the angle between this vector and another in radians
func (v *Vector) Angle(vec *Vector) float64 {
	dot := clamp(v.Dot(vec)/(v.Length()*vec.Length()), -1.0, 1.0)
	return math.Acos(dot)
}

// Dot calculates the dot product of this vector with another
func (v *Vector) Dot(vec *Vector) float64 {
	return v.x*vec.x + v.y*vec.y + v.z*vec.z
}

// Cross calculates the cross product of this vector with another
func (v *Vector) Cross(o *Vector) *Vector {
	return NewVector(
		v.y*o.z-o.y*v.z,
		v.z*o.x-o.z*v.x,
		v.x*o.y-o.x*v.y,
	)
}

// RotateAroundX rotates the vector around the x-axis
func (v *Vector) RotateAroundX(angle float64) *Vector {
	angleCos := math.Cos(angle)
	angleSin := math.Sin(angle)

	newY := angleCos*v.y - angleSin*v.z
	newZ := angleSin*v.y + angleCos*v.z
	return NewVector(v.x, newY, newZ)
}

// RotateAroundY rotates the vector around the y-axis
func (v *Vector) RotateAroundY(angle float64) *Vector {
	angleCos := math.Cos(angle)
	angleSin := math.Sin(angle)

	newX := angleCos*v.x + angleSin*v.z
	newZ := -angleSin*v.x + angleCos*v.z
	return NewVector(newX, v.y, newZ)
}

// RotateAroundZ rotates the vector around the z-axis
func (v *Vector) RotateAroundZ(angle float64) *Vector {
	angleCos := math.Cos(angle)
	angleSin := math.Sin(angle)

	newX := angleCos*v.x - angleSin*v.y
	newY := angleSin*v.x + angleCos*v.y
	return NewVector(newX, newY, v.z)
}

// Rotate rotates the surrounding vector three axes
func (v *Vector) Rotate(angleX, angleY, angleZ float64) *Vector {
	return v.RotateAroundX(angleX).RotateAroundY(angleY).RotateAroundZ(angleZ)
}

// RotateFromView rotates the vector from view angles (yaw and pitch in degrees)
func (v *Vector) RotateFromView(yawDegrees, pitchDegrees float64) *Vector {
	yaw := math.Pi * (-1 * (yawDegrees + 90)) / 180
	pitch := math.Pi * (-pitchDegrees) / 180

	cosYaw := math.Cos(yaw)
	cosPitch := math.Cos(pitch)
	sinYaw := math.Sin(yaw)
	sinPitch := math.Sin(pitch)

	// Z_Axis rotation (Pitch)
	initialX := v.x
	initialY := v.y
	x := initialX*cosPitch - initialY*sinPitch
	y := initialX*sinPitch + initialY*cosPitch

	// Y_Axis rotation (Yaw)
	initialZ := v.z
	initialX = x
	z := initialZ*cosYaw - initialX*sinYaw
	x = initialZ*sinYaw + initialX*cosYaw

	return NewVector(x, y, z)
}

// RotateAroundAxis rotates the vector around a given arbitrary axis in 3D space
func (v *Vector) RotateAroundAxis(axis *Vector, angle float64) *Vector {
	normalizedAxis := axis
	if !axis.IsNormalized() {
		normalizedAxis = axis.Normalize()
	}
	return v.RotateAroundNonUnitAxis(normalizedAxis, angle)
}

// RotateAroundNonUnitAxis rotates the vector around a given arbitrary axis (may not be unit vector)
func (v *Vector) RotateAroundNonUnitAxis(axis *Vector, angle float64) *Vector {
	x, y, z := v.x, v.y, v.z
	x2, y2, z2 := axis.x, axis.y, axis.z
	cosTheta := math.Cos(angle)
	sinTheta := math.Sin(angle)
	dotProduct := v.Dot(axis)

	newX := x2*dotProduct*(1-cosTheta) + x*cosTheta + (-z2*y+y2*z)*sinTheta
	newY := y2*dotProduct*(1-cosTheta) + y*cosTheta + (z2*x-x2*z)*sinTheta
	newZ := z2*dotProduct*(1-cosTheta) + z*cosTheta + (-y2*x+x2*y)*sinTheta

	return NewVector(newX, newY, newZ)
}

// Lerp calculates a linear interpolation between this vector and another
func (v *Vector) Lerp(vec *Vector, alpha float64) *Vector {
	return NewVector(
		v.x+(alpha*(vec.x-v.x)),
		v.y+(alpha*(vec.y-v.y)),
		v.z+(alpha*(vec.z-v.z)),
	)
}

// Interpolate interpolates between this vector and target using the specified interpolation
func (v *Vector) Interpolate(target *Vector, alpha float64, interpolation Interpolation) *Vector {
	return v.Lerp(target, interpolation(alpha))
}

// Utility functions
func globalToBlock(coordinate float64) int {
	return int(math.Floor(coordinate))
}

func globalToChunk(coordinate float64) int {
	return int(math.Floor(coordinate)) >> 4
}

func signum(value float64) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	}
	return 0
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
