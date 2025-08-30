package movable

import (
	"example/hello/src/aabb"
	"example/hello/src/constants"
)

type NonMovableObjectData struct {
	ID     constants.ID
	Bounds aabb.AABB
}

func NewNonMovableObject(id constants.ID, bounds aabb.AABB) NonMovableObjectData {
	return NonMovableObjectData{ID: id, Bounds: bounds}
}
func (o NonMovableObjectData) GetBounds() aabb.AABB {
	return o.Bounds
}
func (o NonMovableObjectData) GetId() constants.ID {
	return o.ID
}
