package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player draw method
func (p *Player) Draw(scr *ebiten.Image) {
	w, h := p.Sprite.Size()
	optionsPlayer := ebiten.DrawImageOptions{} // Translate image to rotate along center
	optionsPlayer.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	optionsPlayer.GeoM.Rotate(float64(p.Rotation+90) * (math.Pi / 180))
	optionsPlayer.GeoM.Scale(0.1, 0.1)
	optionsPlayer.GeoM.Translate(p.PositionX, p.PositionY)
	scr.DrawImage(p.Sprite, &optionsPlayer)
}
