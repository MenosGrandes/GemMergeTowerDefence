package maputil

import (
	"example/hello/src/aabb"
	"example/hello/src/constants"
	"example/hello/src/drawable"
	imageutil "example/hello/src/imageUtil"
	"example/hello/src/movable"
	"example/hello/src/object"
	"example/hello/src/tower"
	"example/hello/src/vec"
	wavestartingpoint "example/hello/src/waveStartingPoint"
	"fmt"
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type StartingObjects struct {
	Orb                *object.Orb
	Towers             []*tower.Tower
	WaveStartingPoints []*wavestartingpoint.WaveStartingPoint
}

var instance *StartingObjects
var once sync.Once

func GetStartingObjects() *StartingObjects {
	once.Do(func() {
		instance = &StartingObjects{}
	})
	return instance
}

const (
	OrbClassName           = "orb"
	StartingPointClassName = "starting_point"
)

func CalculateNearest32x32(pos vec.Vec2F64) vec.Vec2F64 {
	//I must calculate center of 32x32
	fmt.Println(pos)
	cx := float64((int(pos.X/constants.SpriteSize) * constants.SpriteSize) + (constants.SpriteSize / 2))
	cy := float64((int(pos.Y/constants.SpriteSize) * constants.SpriteSize) + (constants.SpriteSize / 2))

	center := vec.NewVec2F64(cx, cy)
	fmt.Println(center)
	return center
}
func LoadStartingObjects(mapData *tiled.Map) {

	neededImages := imageutil.GetNeededImages()
	startingObjects := GetStartingObjects()

	for _, og := range mapData.ObjectGroups {
		fmt.Println(og)
		for _, o := range og.Objects {
			//res2B, _ := json.MarshalIndent(o, "", "    ")
			//fmt.Println(string(res2B))

			//fmt.Println(o)
			//MenosGrandes for some reason the Type is still valid, and it uses Class
			switch o.Type {
			case OrbClassName:
				fmt.Println("Read orb")
				rect2 := image.Rect(320, 32, 352, 64)

				bounds := aabb.NewAABBSprite(CalculateNearest32x32(vec.NewVec2F64(o.X, o.Y)))
				nmo := movable.NewNonMovableObject(1, bounds)
				obj := object.NewOrbPtr(nmo, drawable.NewDrawablePtr(neededImages.PathTileset.SubImage(rect2).(*ebiten.Image)))
				startingObjects.Orb = obj
			case StartingPointClassName:
				fmt.Println("Read StartingPOint")
				rect2 := image.Rect(320, 32, 352, 64)

				bounds := aabb.NewAABBSprite(CalculateNearest32x32(vec.NewVec2F64(o.X, o.Y)))
				nmo := movable.NewNonMovableObject(1, bounds)
				obj := wavestartingpoint.NewWaveStartingPointPtr(nmo, drawable.NewDrawablePtr(neededImages.PathTileset.SubImage(rect2).(*ebiten.Image)))
				startingObjects.WaveStartingPoints = append(startingObjects.WaveStartingPoints, obj)
			}

		}

	}
}
