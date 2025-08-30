package imageutil

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

type NeededImages struct {
	PathTileset      *ebiten.Image
	BacgroundPicture *ebiten.Image
}

var instance *NeededImages
var once sync.Once

func GetNeededImages() *NeededImages {
	once.Do(func() {
		instance = &NeededImages{}
	})
	return instance
}

func GetAllImagesFromMap(mapData *tiled.Map) {
	neededImages := GetNeededImages()
	for _, ts := range mapData.Tilesets {
		if ts.Name == "pathT" {
			neededImages.PathTileset = loadTilesImage(ts)
			break
		}
	}
	if neededImages.PathTileset == nil {
		log.Fatal("No tileset named path")
		return
	}

	neededImages.BacgroundPicture = loadTileFromTiledDirectory(mapData.ImageLayers[0].Image.Source)
}

func loadTileFromTiledDirectory(source string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(filepath.Join("tiled", source))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func loadTilesImage(ts *tiled.Tileset) *ebiten.Image {
	return loadTileFromTiledDirectory(ts.Image.Source)
}
