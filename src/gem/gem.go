package gem

import (
	"example/hello/src/aabb"
	"example/hello/src/atack_data"
	"example/hello/src/constants"
	"example/hello/src/drawable"
	"example/hello/src/enemy"
	"example/hello/src/movable"
	"example/hello/src/vec"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Gem struct {
	nonMovableObjectData movable.NonMovableObjectData
	Drawable             drawable.DrawableI
	gemRange             float32
	atackData            atack_data.AttackData
}

func NewGem(nmod movable.NonMovableObjectData, drawable drawable.DrawableI, ad atack_data.AttackData) Gem {
	return Gem{nonMovableObjectData: nmod, Drawable: drawable, gemRange: 30, atackData: ad}
}
func NewGemPtr(nmod movable.NonMovableObjectData, drawable drawable.DrawableI, ad atack_data.AttackData) *Gem {
	return &Gem{nonMovableObjectData: nmod, Drawable: drawable, gemRange: 30, atackData: ad}
}

// Draw implements GemI.
func (gem *Gem) Draw(screen *ebiten.Image, dio *ebiten.DrawImageOptions) {

	b := gem.GetBounds()
	gem.Drawable.Draw(screen, dio)

	drawable.DebugDrawUtil.DrawAABB(screen,
		gem.GetBounds(),
		color.RGBA{255, 255, 100, 100})
	drawable.DebugDrawUtil.DrawCircle(screen, b.Center.X, b.Center.Y, float64(gem.gemRange), color.Black) //MenosGrandes draw order is f up
}

// GetBounds implements GemI.
func (gem *Gem) GetBounds() aabb.AABB {
	return gem.nonMovableObjectData.GetBounds()

}

// GetId implements GemI.
func (gem *Gem) GetId() constants.ID {
	return gem.nonMovableObjectData.GetId()

}

// IsClicked implements GemI.
func (gem *Gem) IsClicked(mouseClickPos vec.Vec2F64) bool {
	c := gem.GetBounds().ContainsPoint(mouseClickPos)
	if c {
		fmt.Println("Gem Clicked")
	}
	return c
}

// Update implements GemI.
func (gem *Gem) GetRange() float32 {
	return gem.gemRange
}

// Update implements GemI.
func (gem *Gem) Update() {
}

func (gem *Gem) Attack(e enemy.EnemyI) {
	e.OnAttack(gem.atackData)
	//e.SetStats(enemy.Stats{Health: e.GetStats().Health - 1})
	//fmt.Printf("Attack - id[%d], health[%d]\n", e.GetId(), e.GetStats().Health)
}
