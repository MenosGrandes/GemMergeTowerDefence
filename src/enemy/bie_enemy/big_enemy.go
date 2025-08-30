package big_enemy

import (
	"example/hello/src/aabb"
	"example/hello/src/atack_data"
	"example/hello/src/constants"
	"example/hello/src/enemy"
	"example/hello/src/vec"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type BigEnemy struct {
	enemyData enemy.EnemyData
}

// OnAttack implements enemy.EnemyI.
func (b *BigEnemy) OnAttack(ad atack_data.AttackData) {
	b.enemyData.Stats.Health -= ad.Health
	fmt.Println(b.enemyData.Stats)
}

func (b *BigEnemy) SetStats(stats enemy.Stats) {
	b.enemyData.Stats = stats
}

// Draw implements EnemyI.
func (b *BigEnemy) Draw(screen *ebiten.Image, dio *ebiten.DrawImageOptions) {
	b.enemyData.Drawable.Draw(screen, dio)
}

// GetBounds implements EnemyI.
func (b *BigEnemy) GetBounds() aabb.AABB {
	return b.enemyData.MovableObjectData.GetBounds()
}

// GetId implements EnemyI.
func (b *BigEnemy) GetId() constants.ID {
	return b.enemyData.MovableObjectData.GetId()

}

// GetStats implements EnemyI.
func (b *BigEnemy) GetStats() enemy.Stats {
	return b.enemyData.Stats

}

// IsClicked implements EnemyI.
func (b *BigEnemy) IsClicked(mouseClickPos vec.Vec2F64) bool {
	return false
}

// ShouldDie implements EnemyI.
func (b *BigEnemy) ShouldDie() bool {
	return b.GetStats().Health <= 0
}

// Update implements EnemyI.
func (b *BigEnemy) Update() {

}
func (b *BigEnemy) Move(newPosition constants.Position) {
	b.enemyData.MovableObjectData.Move(newPosition)
}
