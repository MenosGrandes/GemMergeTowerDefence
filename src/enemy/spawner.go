package enemy

import (
	"example/hello/src/constants"
	"example/hello/src/drawable"
	imageutil "example/hello/src/imageUtil"
	"example/hello/src/movable"
	wavestartingpoint "example/hello/src/waveStartingPoint"
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemySpawner struct {
	Enemies []EnemyI
}

func (e *EnemySpawner) SpawnEnemy(n int, startingPoint *wavestartingpoint.WaveStartingPoint, factory EnemyFactoryI) {
	//pool := GetEnemyPool()
	ts := imageutil.GetNeededImages().PathTileset
	rect2 := image.Rect(320, 32, 352, 64)
	speed := 5.0
	for range n {
		movable := movable.NewMovableObject(1, startingPoint.GetBounds(), constants.Position(startingPoint.GetBounds().Center), speed)
		drawable := drawable.NewDrawablePtr(ts.SubImage(rect2).(*ebiten.Image)) //MENOSGRANDES Im creating the same item whole time
		enemyData := NewEnemyData(movable, drawable, Stats{Health: 100})
		enemy := factory.CreateEnemy(enemyData)
		e.Enemies = append(e.Enemies, enemy)
	}
}

// func (e *EnemySpawner) DestroyEnemy(obj *Enemy) {
// 	GetEnemyPool().Put(obj)
// }

func (e *EnemySpawner) RemoveDead() {
	filtered := e.Enemies[:0]
	for _, enemy := range e.Enemies {
		if !enemy.ShouldDie() {
			filtered = append(filtered, enemy)
		}
	}
	e.Enemies = filtered
}

var es_instance *EnemySpawner
var es_once sync.Once

func GetEnemySpawner() *EnemySpawner {
	es_once.Do(func() {
		es_instance = &EnemySpawner{}
	})
	return es_instance
}
