package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Drawable struct {
	sprite *ebiten.Image
}

func (e Drawable) Draw(screen *ebiten.Image, dio *ebiten.DrawImageOptions) {
	screen.DrawImage(e.sprite, dio)
}
func NewDrawablePtr(bg *ebiten.Image) *Drawable {
	return &Drawable{bg}
}
func NewDrawableEmptyPtr() *Drawable {
	return &Drawable{nil}
}
