package aabb

import (
	"example/hello/src/constants"
	"example/hello/src/vec"
	"math"
)

type AABB struct {
	Center                vec.Vec2F64
	HalfWidth, HalfHeight float64
}

func NewAABB(center vec.Vec2F64, halfWidth, halfHeight float64) AABB {
	return AABB{Center: center, HalfWidth: halfWidth, HalfHeight: halfHeight}
}
func NewAABBSprite(center vec.Vec2F64) AABB {
	return AABB{Center: center, HalfWidth: constants.SpriteSize / 2, HalfHeight: constants.SpriteSize / 2}
}
func (a AABB) Intersects(b AABB) bool {
	return math.Abs(a.Center.X-b.Center.X) <= (a.HalfWidth+b.HalfWidth) &&
		math.Abs(a.Center.Y-b.Center.Y) <= (a.HalfHeight+b.HalfHeight)
}

func (a AABB) Contains(b AABB) bool {
	return math.Abs(b.Center.X-a.Center.X)+b.HalfWidth <= a.HalfWidth &&
		math.Abs(b.Center.Y-a.Center.Y)+b.HalfHeight <= a.HalfHeight
}
func (a AABB) ContainsPoint(point vec.Vec2F64) bool {
	return point.X >= a.Center.X-a.HalfWidth && point.X <= a.Center.X+a.HalfWidth &&
		point.Y >= a.Center.Y-a.HalfHeight && point.Y <= a.Center.Y+a.HalfHeight
}
