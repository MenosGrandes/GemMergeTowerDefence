package obj_i

import (
	"example/hello/src/aabb"
	"example/hello/src/clickable"
	"example/hello/src/constants"
	"example/hello/src/drawable"
	"example/hello/src/updatable"
)

type ObjectI interface {
	clickable.MouseClickableI
	drawable.DrawableI
	updatable.UpdatableI
	GetBounds() aabb.AABB
	GetId() constants.ID
}
