package gem

import (
	"example/hello/src/enemy"
	"example/hello/src/object/obj_i"
)

type GemI interface {
	obj_i.ObjectI
	Attack(e enemy.EnemyI)
	GetRange() float32
}
