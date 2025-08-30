package enemy

import (
	"example/hello/src/drawable"
	"example/hello/src/movable"
)

type EnemyData struct {
	MovableObjectData movable.MovableObjectData
	Drawable          drawable.DrawableI
	Stats             Stats
}

func NewEnemyData(mod movable.MovableObjectData, drawable drawable.DrawableI, stats Stats) EnemyData {
	return EnemyData{MovableObjectData: mod, Drawable: drawable, Stats: stats}
}
