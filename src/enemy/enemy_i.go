package enemy

import (
	"example/hello/src/atack_data"
	"example/hello/src/movable"
	"example/hello/src/object/obj_i"
)

type Stats struct {
	Health int64
}
type EnemyI interface {
	obj_i.ObjectI
	movable.MovableI
	GetStats() Stats
	ShouldDie() bool
	SetStats(stats Stats)
	OnAttack(atack_data.AttackData)
}
