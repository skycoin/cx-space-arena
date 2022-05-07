package asteroid

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Asteroid draw method
func (a *Asteroid) Draw(scr *ebiten.Image) {
	w, h := a.Sprite.Size()
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	options.GeoM.Rotate(float64(a.Rotation) * (math.Pi / 180))
	options.GeoM.Scale(0.05*a.Mass, 0.05*a.Mass) // Size between 0.1 and 0.2
	options.GeoM.Translate(a.PositionX, a.PositionY)
	scr.DrawImage(a.Sprite, &options)
}
