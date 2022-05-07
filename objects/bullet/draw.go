package bullet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (b *Bullet) Draw(scr *ebiten.Image) {
	w, h := b.Sprite.Size()
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	options.GeoM.Rotate(float64(b.Angle) * (math.Pi / 180))
	options.GeoM.Scale(0.1, 0.1)
	options.GeoM.Translate(b.PositionX, b.PositionY)
	scr.DrawImage(b.Sprite, &options)
}
