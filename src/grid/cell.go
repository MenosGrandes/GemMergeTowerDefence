package grid

import (
	"example/hello/src/constants"
	"example/hello/src/drawable"
	"example/hello/src/object/obj_i"
	"example/hello/src/vec"
)

// MenosGrandes this should implement Drawable
type Cell struct {
	ID        constants.ID
	Drawable  drawable.DrawableI
	Object    obj_i.ObjectI //MENOSGRANDES should be another type. Like GemPlacable? As this can contain only Tower/Trap/Block
	Direction vec.Vec2F64
}

func NewCell(ID constants.ID) Cell {
	return Cell{ID, nil, nil, vec.NewVec2F64(0, 0)}
}
func NewCellPtr(ID constants.ID) *Cell {
	return &Cell{ID, nil, nil, vec.NewVec2F64(0, 0)}
}
func NewCellDrawable(ID constants.ID, drawable drawable.DrawableI) Cell {
	return Cell{ID, drawable, nil, vec.NewVec2F64(0, 0)}
}
func NewCellDrawablePtr(ID constants.ID, drawable drawable.DrawableI) *Cell {
	return &Cell{ID, drawable, nil, vec.NewVec2F64(0, 0)}
}
func (e *Cell) AddObject(object obj_i.ObjectI) {
	e.Object = object
}

func (e *Cell) IsClicked(mouseClickPos vec.Vec2F64) bool {
	if e.Object != nil {
		return e.Object.IsClicked(mouseClickPos)

	}

	return false
}
