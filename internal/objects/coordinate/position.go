package coordinate

import (
	"Veloce/internal/objects/block"
	"math"
)

// Position represents an immutable position containing coordinates and a view (yaw/pitch)
type Position struct {
	x     float64
	y     float64
	z     float64
	yaw   float32
	pitch float32
}

// Common position constants

var (
	PosZero = NewPos(0, 0, 0)
)

// NewPos creates a new position with coordinates and default view angles (0, 0)
func NewPos(x, y, z float64) *Position {
	return &Position{
		x:     x,
		y:     y,
		z:     z,
		yaw:   0,
		pitch: 0,
	}
}

// NewPosWithView creates a new position with coordinates and view angles
func NewPosWithView(x, y, z float64, yaw, pitch float32) *Position {
	return &Position{
		x:     x,
		y:     y,
		z:     z,
		yaw:   fixYaw(yaw),
		pitch: pitch,
	}
}

// NewPosFromPoint creates a position from a Point with default view angles
func NewPosFromPoint(point Point) *Position {
	return NewPos(point.X(), point.Y(), point.Z())
}

// NewPosFromPointWithView creates a position from a Point with specified view angles
func NewPosFromPointWithView(point Point, yaw, pitch float32) *Position {
	return NewPosWithView(point.X(), point.Y(), point.Z(), yaw, pitch)
}

// PosFromPoint converts a Point into a Position, casting if possible or creating new
func PosFromPoint(point Point) *Position {
	if pos, ok := point.(*Position); ok {
		return pos
	}
	return NewPosFromPoint(point)
}

// Yaw returns the yaw angle
func (p *Position) Yaw() float32 {
	return p.yaw
}

// Pitch returns the pitch angle
func (p *Position) Pitch() float32 {
	return p.pitch
}

// Point interface implementation

func (p *Position) X() float64 {
	return p.x
}

func (p *Position) Y() float64 {
	return p.y
}

func (p *Position) Z() float64 {
	return p.z
}

func (p *Position) BlockX() int {
	return globalToBlock(p.x)
}

func (p *Position) BlockY() int {
	return globalToBlock(p.y)
}

func (p *Position) BlockZ() int {
	return globalToBlock(p.z)
}

func (p *Position) ChunkX() int {
	return globalToChunk(p.x)
}

func (p *Position) Section() int {
	return globalToChunk(p.y)
}

func (p *Position) ChunkZ() int {
	return globalToChunk(p.z)
}

func (p *Position) WithX(x float64) Point {
	return NewPosWithView(x, p.y, p.z, p.yaw, p.pitch)
}

func (p *Position) WithY(y float64) Point {
	return NewPosWithView(p.x, y, p.z, p.yaw, p.pitch)
}

func (p *Position) WithZ(z float64) Point {
	return NewPosWithView(p.x, p.y, z, p.yaw, p.pitch)
}

func (p *Position) Add(x, y, z float64) Point {
	return NewPosWithView(p.x+x, p.y+y, p.z+z, p.yaw, p.pitch)
}

func (p *Position) AddPoint(point Point) Point {
	return p.Add(point.X(), point.Y(), point.Z())
}

func (p *Position) AddValue(value float64) Point {
	return p.Add(value, value, value)
}

func (p *Position) Sub(x, y, z float64) Point {
	return NewPosWithView(p.x-x, p.y-y, p.z-z, p.yaw, p.pitch)
}

func (p *Position) SubPoint(point Point) Point {
	return p.Sub(point.X(), point.Y(), point.Z())
}

func (p *Position) SubValue(value float64) Point {
	return p.Sub(value, value, value)
}

func (p *Position) Mul(x, y, z float64) Point {
	return NewPosWithView(p.x*x, p.y*y, p.z*z, p.yaw, p.pitch)
}

func (p *Position) MulPoint(point Point) Point {
	return p.Mul(point.X(), point.Y(), point.Z())
}

func (p *Position) MulValue(value float64) Point {
	return p.Mul(value, value, value)
}

func (p *Position) Div(x, y, z float64) Point {
	return NewPosWithView(p.x/x, p.y/y, p.z/z, p.yaw, p.pitch)
}

func (p *Position) DivPoint(point Point) Point {
	return p.Div(point.X(), point.Y(), point.Z())
}

func (p *Position) DivValue(value float64) Point {
	return p.Div(value, value, value)
}

func (p *Position) Relative(face block.Face) Point {
	switch face {
	case block.Bottom:
		return p.Sub(0, 1, 0)
	case block.Top:
		return p.Add(0, 1, 0)
	case block.North:
		return p.Sub(0, 0, 1)
	case block.South:
		return p.Add(0, 0, 1)
	case block.West:
		return p.Sub(1, 0, 0)
	case block.East:
		return p.Add(1, 0, 0)
	default:
		return p
	}
}

func (p *Position) DistanceSquared(x, y, z float64) float64 {
	dx := p.x - x
	dy := p.y - y
	dz := p.z - z
	return dx*dx + dy*dy + dz*dz
}

func (p *Position) DistanceSquaredToPoint(point Point) float64 {
	return p.DistanceSquared(point.X(), point.Y(), point.Z())
}

func (p *Position) Distance(x, y, z float64) float64 {
	return math.Sqrt(p.DistanceSquared(x, y, z))
}

func (p *Position) DistanceToPoint(point Point) float64 {
	return p.Distance(point.X(), point.Y(), point.Z())
}

func (p *Position) SamePoint(x, y, z float64) bool {
	return p.x == x && p.y == y && p.z == z
}

func (p *Position) SamePointAs(point Point) bool {
	return p.SamePoint(point.X(), point.Y(), point.Z())
}

func (p *Position) IsZero() bool {
	return p.x == 0 && p.y == 0 && p.z == 0
}

func (p *Position) SameChunk(point Point) bool {
	return p.ChunkX() == point.ChunkX() && p.ChunkZ() == point.ChunkZ()
}

func (p *Position) SameBlock(blockX, blockY, blockZ int) bool {
	return p.BlockX() == blockX && p.BlockY() == blockY && p.BlockZ() == blockZ
}

func (p *Position) SameBlockAs(point Point) bool {
	return p.SameBlock(point.BlockX(), point.BlockY(), point.BlockZ())
}

// Position-specific methods

// WithCoord creates a new position with new coordinates, preserving view angles
func (p *Position) WithCoord(x, y, z float64) *Position {
	return NewPosWithView(x, y, z, p.yaw, p.pitch)
}

// WithCoordPoint creates a new position with coordinates from a point, preserving view angles
func (p *Position) WithCoordPoint(point Point) *Position {
	return p.WithCoord(point.X(), point.Y(), point.Z())
}

// WithView creates a new position with new view angles, preserving coordinates
func (p *Position) WithView(yaw, pitch float32) *Position {
	return NewPosWithView(p.x, p.y, p.z, yaw, pitch)
}

// WithViewFromPos creates a new position with view angles from another position
func (p *Position) WithViewFromPos(pos *Position) *Position {
	return p.WithView(pos.yaw, pos.pitch)
}

// WithDirection sets the yaw and pitch to point in the direction of the given point
func (p *Position) WithDirection(point Point) *Position {
	x := point.X()
	z := point.Z()

	if x == 0 && z == 0 {
		pitch := float32(90)
		if point.Y() > 0 {
			pitch = -90
		}
		return p.WithPitch(pitch)
	}

	theta := math.Atan2(-x, z)
	xz := math.Sqrt(x*x + z*z)
	_2PI := 2 * math.Pi

	yaw := float32(math.Mod(theta+_2PI, _2PI) * 180.0 / math.Pi)
	pitch := float32(math.Atan(-point.Y()/xz) * 180.0 / math.Pi)

	return p.WithView(yaw, pitch)
}

// WithYaw creates a new position with a new yaw angle
func (p *Position) WithYaw(yaw float32) *Position {
	return NewPosWithView(p.x, p.y, p.z, yaw, p.pitch)
}

// WithPitch creates a new position with a new pitch angle
func (p *Position) WithPitch(pitch float32) *Position {
	return NewPosWithView(p.x, p.y, p.z, p.yaw, pitch)
}

// WithLookAt creates a new position looking at the specified point
func (p *Position) WithLookAt(point Point) *Position {
	if p.SamePointAs(point) {
		return p
	}

	delta := VectorFromPoint(point.SubPoint(p)).Normalize()
	yaw := getLookYaw(delta.X(), delta.Z())
	pitch := getLookPitch(delta.X(), delta.Y(), delta.Z())

	return p.WithView(yaw, pitch)
}

// SameView checks if two positions have the same view (yaw/pitch)
func (p *Position) SameView(pos *Position) bool {
	return p.SameViewAngles(pos.yaw, pos.pitch)
}

// SameViewAngles checks if the position has the same view angles as specified
func (p *Position) SameViewAngles(yaw, pitch float32) bool {
	return p.yaw == yaw && p.pitch == pitch
}

// Direction gets a unit-vector pointing in the direction this position is facing
func (p *Position) Direction() *Vector {
	rotX := float64(p.yaw)
	rotY := float64(p.pitch)
	xz := math.Cos(math.Pi * rotY / 180.0)

	return NewVector(
		-xz*math.Sin(math.Pi*rotX/180.0),
		-math.Sin(math.Pi*rotY/180.0),
		xz*math.Cos(math.Pi*rotX/180.0),
	)
}

// AsVec converts this position to a vector (losing view angles)
func (p *Position) AsVec() *Vector {
	return NewVector(p.x, p.y, p.z)
}

// Apply applies an operator function to this position
func (p *Position) Apply(operator func(x, y, z float64, yaw, pitch float32) *Position) *Position {
	return operator(p.x, p.y, p.z, p.yaw, p.pitch)
}

// Utility functions

// fixYaw normalizes a yaw value to be between -180 and 180 degrees
func fixYaw(yaw float32) float32 {
	yaw = float32(math.Mod(float64(yaw), 360))
	if yaw < -180.0 {
		yaw += 360.0
	} else if yaw > 180.0 {
		yaw -= 360.0
	}
	return yaw
}

// getLookYaw calculates yaw from direction components
func getLookYaw(x, z float64) float32 {
	yaw := math.Atan2(-x, z)
	return float32(yaw * 180.0 / math.Pi)
}

// getLookPitch calculates pitch from direction components
func getLookPitch(x, y, z float64) float32 {
	xz := math.Sqrt(x*x + z*z)
	pitch := math.Atan2(-y, xz)
	return float32(pitch * 180.0 / math.Pi)
}
