package object

import (
	"example/hello/src/aabb"
	"example/hello/src/constants"
	"example/hello/src/drawable"
	"example/hello/src/movable"
	"example/hello/src/vec"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Orb struct {
	nonMovableObjectData movable.NonMovableObjectData
	Drawable             drawable.DrawableI
}

func (e *Orb) GetBounds() aabb.AABB {
	return e.nonMovableObjectData.Bounds
}
func (e *Orb) GetId() constants.ID {
	return e.nonMovableObjectData.ID
}
func (e *Orb) Update() {

}
func (e *Orb) Draw(screen *ebiten.Image, dio *ebiten.DrawImageOptions) {
	e.Drawable.Draw(screen, dio)

	drawable.DebugDrawUtil.DrawAABB(screen,
		e.GetBounds(),
		color.RGBA{255, 100, 100, 100})

}
func NewOrbPtr(nonMovableObjectData movable.NonMovableObjectData, drawable *drawable.Drawable) *Orb {
	return &Orb{nonMovableObjectData, drawable}
}
func (e *Orb) IsClicked(mouseClickPos vec.Vec2F64) bool {
	c := e.GetBounds().ContainsPoint(mouseClickPos)
	if c {
		fmt.Println("Orb Clicked")
	}
	return c
}
