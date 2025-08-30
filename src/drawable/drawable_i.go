package drawable

import "github.com/hajimehoshi/ebiten/v2"

type DrawableI interface {
	Draw(screen *ebiten.Image, dio *ebiten.DrawImageOptions)
}
