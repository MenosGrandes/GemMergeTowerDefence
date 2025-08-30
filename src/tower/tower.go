package tower

import (
	"example/hello/src/aabb"
	"example/hello/src/constants"
	"example/hello/src/drawable"
	"example/hello/src/enemy"
	"example/hello/src/gem"
	"example/hello/src/movable"
	"example/hello/src/vec"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tower struct {
	nonMovableObjectData movable.NonMovableObjectData
	Drawable             drawable.DrawableI
	Gem                  gem.GemI
}

func (e Tower) GetBounds() aabb.AABB {
	return e.nonMovableObjectData.Bounds
}
func (e Tower) GetId() constants.ID {
	return e.nonMovableObjectData.ID
}
func (e *Tower) Update() {
	if e.Gem != nil {
		enemy := getNearestEnemy()
		if enemy != nil {
			e.Gem.Attack(enemy)
		}
	}
}

func getNearestEnemy() enemy.EnemyI {
	es := enemy.GetEnemySpawner().Enemies
	e := len(es)
	if e != 0 {
		return es[0]
	}
	return nil
}

func (e Tower) Draw(screen *ebiten.Image, dio *ebiten.DrawImageOptions) {
	e.Drawable.Draw(screen, dio)
	drawable.DebugDrawUtil.DrawAABB(screen,
		e.GetBounds(),
		color.RGBA{0, 255, 100, 255})

	if e.Gem != nil {
		e.Gem.Draw(screen, dio)
	}
}
func NewTowerPtr(nonMovableObjectData movable.NonMovableObjectData, drawable drawable.DrawableI) *Tower {
	return &Tower{nonMovableObjectData, drawable, nil}
}
func (e Tower) IsClicked(mouseClickPos vec.Vec2F64) bool {
	c := e.GetBounds().ContainsPoint(mouseClickPos)
	if c {
		fmt.Println("TOWER Clicked")
	}
	return c
}

func (e *Tower) AddGem(gem gem.GemI) {
	e.Gem = gem
}
