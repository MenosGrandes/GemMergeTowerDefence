package wavestartingpoint

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

type WaveStartingPoint struct {
	nonMovableObjectData movable.NonMovableObjectData
	Drawable             drawable.DrawableI
}

func (e *WaveStartingPoint) GetBounds() aabb.AABB {
	return e.nonMovableObjectData.Bounds
}
func (e *WaveStartingPoint) GetId() constants.ID {
	return e.nonMovableObjectData.ID
}

func (e *WaveStartingPoint) Draw(screen *ebiten.Image, dio *ebiten.DrawImageOptions) {
	e.Drawable.Draw(screen, dio)
	drawable.DebugDrawUtil.DrawAABB(screen,
		e.GetBounds(),
		color.RGBA{100, 100, 255, 100})
}
func NewWaveStartingPointPtr(nonMovableObjectData movable.NonMovableObjectData, drawable *drawable.Drawable) *WaveStartingPoint {
	return &WaveStartingPoint{nonMovableObjectData, drawable}
}
func (e *WaveStartingPoint) IsClicked(mouseClickPos vec.Vec2F64) bool {
	c := e.GetBounds().ContainsPoint(mouseClickPos)
	if c {
		fmt.Println("WaveStartingPoint Clicked")
	}
	return c
}
