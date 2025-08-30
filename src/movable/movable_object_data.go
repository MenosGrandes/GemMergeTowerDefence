package movable

import (
	"example/hello/src/aabb"
	"example/hello/src/constants"
	"example/hello/src/vec"
)

type MovableObjectData struct {
	ID       constants.ID
	Bounds   aabb.AABB
	Position constants.Position
	Speed    float64
}

func NewMovableObjectEmpty(id constants.ID) MovableObjectData {
	return MovableObjectData{ID: id, Bounds: aabb.AABB{}, Position: constants.Position{}}
}
func NewMovableObject(id constants.ID, bounds aabb.AABB, position constants.Position, speed float64) MovableObjectData {
	return MovableObjectData{ID: id, Bounds: bounds, Position: position, Speed: speed}
}
func (o MovableObjectData) GetBounds() aabb.AABB {
	return o.Bounds
}
func (o MovableObjectData) GetId() constants.ID {
	return o.ID
}
func (o *MovableObjectData) recalculateBounds() {
	o.Bounds = aabb.NewAABBSprite(vec.Vec2F64(o.Position))
}

func (o *MovableObjectData) Move(newPosition constants.Position) {
	o.Position = constants.Position(vec.MoveTowards(vec.Vec2[float64](o.Position), vec.Vec2[float64](newPosition), o.Speed))
	o.recalculateBounds()
}
