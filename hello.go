package main

import (
	"example/hello/src/aabb"
	"example/hello/src/atack_data"
	"example/hello/src/constants"
	"example/hello/src/drawable"
	"example/hello/src/enemy"
	big_enemy "example/hello/src/enemy/bie_enemy"
	"example/hello/src/gem"
	"example/hello/src/grid"
	imageutil "example/hello/src/imageUtil"
	maputil "example/hello/src/mapUtil"
	"example/hello/src/movable"
	"example/hello/src/path"
	"example/hello/src/tower"
	"example/hello/src/vec"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lafriks/go-tiled"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	NumObjects   = 30
	MaxObjects   = 10
	MaxLevels    = 5
)
const mapPath = "tiled/map2.tmx" // Path to your Tiled Map.

type Game struct {
	// Quadtree *quadtree.Quadtree
	fps      float64
	lastTime time.Time
	frames   int
	nextID   int
	mapData  *tiled.Map
	grid     grid.Grid
	ef       enemy.EnemyFactoryI
}

func (g *Game) Update() error {

	objects := maputil.GetStartingObjects()
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		enemy.GetEnemySpawner().SpawnEnemy(1, objects.WaveStartingPoints[0], g.ef)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		path.CalcPath(&g.grid, *g.grid.Cells[0], *g.grid.Cells[len(g.grid.Cells)-1])
	}

	//MenosGrandes write this function so that I can get out of it if something was clicked
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		mousePos := vec.NewVec2F64(float64(x), float64(y))
		cellId := maputil.CalculateNearest32x32(mousePos)
		cellID := int(int(cellId.Y/constants.SpriteSize)*int(g.grid.Size.X) + int(cellId.X/constants.SpriteSize))

		if cellID >= 0 && cellID < int(g.grid.Size.X*g.grid.Size.Y) {
			cell := g.grid.Cells[cellID]
			if cell != nil {
				if cell.Object == nil {
					fmt.Print("ADD TOWER")
					rect2 := image.Rect(320, 32, 352, 64)

					x2 := float64((int(cellID) % int(20)) * constants.SpriteSize)
					y2 := float64((int(cellID) / int(20)) * constants.SpriteSize)
					a := imageutil.GetNeededImages().PathTileset.SubImage(rect2).(*ebiten.Image)
					nmo := movable.NewNonMovableObject(1, aabb.NewAABB(vec.NewVec2F64(x2+(constants.SpriteSize/2), y2+(constants.SpriteSize/2)), (constants.SpriteSize/2), (constants.SpriteSize/2)))
					tower := tower.NewTowerPtr(nmo, drawable.NewDrawablePtr(a))

					cell.AddObject(tower)

				} else {
					fmt.Print("ADD GEM")

					//If there is an Object, it must be Tower.. as there is no other implementation, right?
					t := cell.Object.(*tower.Tower)
					rect2 := image.Rect(0, 96, 32, 96+32)

					x2 := float64((int(cellID) % int(20)) * constants.SpriteSize)
					y2 := float64((int(cellID) / int(20)) * constants.SpriteSize)
					a := imageutil.GetNeededImages().PathTileset.SubImage(rect2).(*ebiten.Image)
					nmo := movable.NewNonMovableObject(1, aabb.NewAABB(vec.NewVec2F64(x2+(constants.SpriteSize/2), y2+(constants.SpriteSize/2)), (constants.SpriteSize/2), (constants.SpriteSize/2)))
					gem := gem.NewGemPtr(nmo, drawable.NewDrawablePtr(a), atack_data.NewAttackData(1, 0))

					t.AddGem(gem)

				}

			}
		}
		for _, c := range g.grid.Cells {
			if c.IsClicked(mousePos) {
				fmt.Println("clicked :", c.ID)
				break
			}
		}
		{
			objects.Orb.IsClicked(mousePos)
		}
		{
			for _, e := range enemy.GetEnemySpawner().Enemies {
				e.IsClicked(mousePos)
			}
		}
	}

	for _, e := range enemy.GetEnemySpawner().Enemies {
		e.Move(constants.Position(objects.Orb.GetBounds().Center))
	}
	//
	// UPDATE
	for _, e := range enemy.GetEnemySpawner().Enemies {
		e.Update()
	}
	enemy.GetEnemySpawner().RemoveDead()
	for _, c := range g.grid.Cells {
		if c.Object != nil {
			c.Object.Update()
		}
	}
	g.frames++
	now := time.Now()
	if now.Sub(g.lastTime) >= time.Second {
		g.fps = float64(g.frames) / now.Sub(g.lastTime).Seconds()
		g.frames = 0
		g.lastTime = now
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//neededImages := imageutil.GetNeededImages()
	//startingObjects := maputil.GetStartingObjects()

	screen.Fill(color.RGBA{0, 0, 0, 255})
	{
		//op := &ebiten.DrawImageOptions{}
		//screen.DrawImage(neededImages.BacgroundPicture, op)
	}
	//MenosGrandes what is happening here?
	/*
		for i, img := range g.grid.Cells {
			if img.Drawable != nil {
				x := (i % int(g.grid.Size.X)) * constants.SpriteSize
				y := (i / int(g.grid.Size.X)) * constants.SpriteSize

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x), float64(y))
				//screen.DrawImage(img.Drawable.Sprite, op)
				img.Drawable.Draw(screen, op)
				if img.Object != nil {
					img.Object.Draw(screen, op)
				}
			} else {
				if img.Object != nil {
					x := (i % int(g.grid.Size.X)) * constants.SpriteSize
					y := (i / int(g.grid.Size.X)) * constants.SpriteSize
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x), float64(y))
					img.Object.Draw(screen, op)
				}
			}
		}
		{
			op := &ebiten.DrawImageOptions{}
			orb := startingObjects.Orb
			op.GeoM.Translate(float64(orb.GetBounds().Center.X-(constants.SpriteSize/2)), float64(orb.GetBounds().Center.Y-(constants.SpriteSize/2)))
			startingObjects.Orb.Draw(screen, op)
		}
		{
			for _, t := range startingObjects.Towers {
				op := &ebiten.DrawImageOptions{}

				op.GeoM.Translate(float64(t.GetBounds().Center.X-(constants.SpriteSize/2)), float64(t.GetBounds().Center.Y-(constants.SpriteSize/2)))
				t.Draw(screen, op)
			}
		}
		{
			for _, t := range startingObjects.WaveStartingPoints {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(t.GetBounds().Center.X-(constants.SpriteSize/2)), float64(t.GetBounds().Center.Y-(constants.SpriteSize/2)))
				t.Draw(screen, op)
			}
		}

		{
			for _, t := range enemy.GetEnemySpawner().Enemies {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(t.GetBounds().Center.X-(constants.SpriteSize/2)), float64(t.GetBounds().Center.Y-(constants.SpriteSize/2)))
				t.Draw(screen, op)
			}
		}
		// g.drawQuadtree(screen, g.Quadtree)
		// for _, obj := range g.Objects {
		// 	DrawRectPolygon(screen, obj.Bounds.Center.X-obj.Bounds.HalfWidth, obj.Bounds.Center.Y-obj.Bounds.HalfHeight, obj.Bounds.HalfWidth*2, obj.Bounds.HalfHeight*2, color.RGBA{255, 180, 100, 255})
		// 	var nearby []object.Object
		// 	g.Quadtree.Retrieve(&nearby, obj)
		// 	for _, other := range nearby {
		// 		if obj.ID != other.ID && obj.Bounds.Intersects(other.Bounds) {
		// 			egitenutil.DrawRectPolygon(screen, obj.Bounds.Center.X-1, obj.Bounds.Center.Y-1, 4, 4, color.RGBA{255, 0, 0, 255})
		// 			break
		// 		}
		// 	}
		// }
	*/
	g.grid.DrawFlowField(screen)
	//drawable.DrawSolidArrow(screen, 200, 200, 0, -1, 50, 0xFF00FF00) // green arrow ↑
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", g.fps))

}

// func (g *Game) drawQuadtree(screen *ebiten.Image, qt *quadtree.Quadtree) {
// egitenutil.DrawRect2(screen,
// 	qt.Bounds.Center.X-qt.Bounds.HalfWidth,
// 	qt.Bounds.Center.Y-qt.Bounds.HalfHeight,
// 	qt.Bounds.HalfWidth*2,
// 	qt.Bounds.HalfHeight*2,
// 	color.RGBA{255, 100, 100, 100})
// 	for _, node := range qt.Nodes {
// 		if node != nil {
// 			g.drawQuadtree(screen, node)
// 		}
// 	}
// }

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	// Parse .tmx file.
	mapData, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	//res2B, _ := json.MarshalIndent(mapData, "", "    ")
	//fmt.Println(string(res2B))

	gridE := grid.NewGridEmpty(vec.NewVec2UI16(uint16(mapData.Width), uint16(mapData.Height)))
	//fmt.Println(len(gridE.Cells))
	//rand.Seed(time.Now().UnixNano())
	// for i := 0; i < NumObjects; i++ {
	// 	size := 1 + rand.Float64()*5
	// 	obj := object.Object{
	// 		ID: i,
	// 		Bounds: aabb.AABB{
	// 			Center:     vec.Vec2F64(rand.Float64()*ScreenWidth, rand.Float64()*ScreenHeight),
	// 			HalfWidth:  size,
	// 			HalfHeight: size,
	// 		},
	// 		Velocity: vec.Vec2F64(rand.Float64()*2-1, rand.Float64()*2-1),
	// 	}
	// 	objects = append(objects, obj)
	// }
	// qt := quadtree.NewQuadtree(0, aabb.AABB{Center: vec.Vec2F64(ScreenWidth/2, ScreenHeight/2), HalfWidth: ScreenWidth / 2, HalfHeight: ScreenHeight / 2})

	imageutil.GetAllImagesFromMap(mapData)
	neededImages := imageutil.GetNeededImages()
	for _, layer := range mapData.Layers {
		if layer.Tiles == nil {
			continue
		}
		if layer.Name == "path" {
			for y := 0; y < mapData.Height; y++ {
				for x := 0; x < mapData.Width; x++ {
					id := constants.ID(y*mapData.Width + x)
					tile := layer.Tiles[id]
					c := grid.NewCellPtr(id)
					if tile == nil {
						gridE.Cells = append(gridE.Cells, c)

						continue
					}
					if tile.Tileset == nil {
						gridE.Cells = append(gridE.Cells, c)

						continue
					}
					if tile.Tileset.Name == "pathT" {
						rect := tile.GetTileRect()
						c = grid.NewCellDrawablePtr(id, drawable.NewDrawablePtr(neededImages.PathTileset.SubImage(rect).(*ebiten.Image)))

						// if rand.IntN(100) < 10 {
						//  rect2 := image.Rect(320, 32, 352, 64)

						// 	x2 := float64((int(id) % int(gridE.Size.X)) * constants.SpriteSize)
						// 	y2 := float64((int(id) / int(gridE.Size.X)) * constants.SpriteSize)

						// 	nmo := object.NewNonMovableObject(1, aabb.NewAABB(vec.NewVec2F64(x2+(constants.SpriteSize/2), y2+(constants.SpriteSize/2)), (constants.SpriteSize/2), (constants.SpriteSize/2)))
						// 	tower := object.NewTowerPtr(nmo, drawable.NewDrawablePtr(neededImages.PathTileset.SubImage(rect2).(*ebiten.Image)))
						// 	c.AddObject(tower)
						// }
					}
					gridE.Cells = append(gridE.Cells, c)

				}
			}
		}

	}
	maputil.LoadStartingObjects(mapData)
	ebiten.SetTPS(60)
	ef := big_enemy.BigEnemyFactory{}
	g := &Game{lastTime: time.Now(), nextID: NumObjects, mapData: mapData, grid: gridE, ef: &ef}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("TD")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
