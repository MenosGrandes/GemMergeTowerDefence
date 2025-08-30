package vec

import "math"

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Vec2[T Numeric] struct {
	X, Y T
}

type Vec2F64 Vec2[float64]
type Vec2I8 Vec2[int8]
type Vec2I16 Vec2[int16]
type Vec2UI16 Vec2[uint16]
type Vec2Int Vec2[int]

func NewEmptyVec[T Numeric]() Vec2[T] {
	return Vec2[T]{}
}

func NewVec2F64(x, y float64) Vec2F64 {
	return Vec2F64{X: x, Y: y}
}
func NewVec2I8(x, y int8) Vec2I8 {
	return Vec2I8{X: x, Y: y}
}
func NewVec2I16(x, y int16) Vec2I16 {
	return Vec2I16{X: x, Y: y}
}
func NewVec2UI16(x, y uint16) Vec2UI16 {
	return Vec2UI16{X: x, Y: y}
}
func NewVec2Int(x, y int) Vec2Int {
	return Vec2Int{X: x, Y: y}
}

// Add two vectors
func (a Vec2[T]) Add(b Vec2[T]) Vec2[T] {
	return Vec2[T]{a.X + b.X, a.Y + b.Y}
}

// Subtract two vectors
func (a Vec2[T]) Sub(b Vec2[T]) Vec2[T] {
	return Vec2[T]{a.X - b.X, a.Y - b.Y}
}

// Scale vector by a scalar
func (v Vec2[T]) Scale(s T) Vec2[T] {
	return Vec2[T]{v.X * s, v.Y * s}
}

// Length (returns float64 for any numeric type)
func (v Vec2[T]) Length() float64 {
	return math.Hypot(float64(v.X), float64(v.Y))
}

// Normalize (returns Vec2[float64] for any numeric type)
func (v Vec2[T]) Normalize() Vec2[float64] {
	len := v.Length()
	if len == 0 {
		return Vec2[float64]{0, 0}
	}
	return Vec2[float64]{float64(v.X) / len, float64(v.Y) / len}
}

func MoveTowards[T Numeric](current, target Vec2[T], speed float64) Vec2[float64] {
	dir := Vec2[float64]{float64(target.X - current.X), float64(target.Y - current.Y)}
	dist := math.Hypot(dir.X, dir.Y)

	if dist <= speed || dist == 0 {
		return Vec2[float64]{float64(target.X), float64(target.Y)}
	}

	move := Vec2[float64]{dir.X * speed / dist, dir.Y * speed / dist}
	return Vec2[float64]{float64(current.X) + move.X, float64(current.Y) + move.Y}
}

func LerpVec2[T float32 | float64](a, b Vec2[T], t T) Vec2[T] {
	return Vec2[T]{
		X: a.X + (b.X-a.X)*t,
		Y: a.Y + (b.Y-a.Y)*t,
	}
}
