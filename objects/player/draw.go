package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player draw method
func (p *Player) Draw(scr *ebiten.Image) {
	w, h := p.Sprite.Size()
	options := ebiten.DrawImageOptions{} // Translate image to rotate along center
	options.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	options.GeoM.Rotate(float64(p.Rotation+90) * (math.Pi / 180))
	options.GeoM.Scale(0.1, 0.1)
	options.GeoM.Translate(p.PositionX, p.PositionY)
	scr.DrawImage(p.Sprite, &options)
}
