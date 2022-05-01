package asteroid

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/assets"
)

func (a *Asteroid) Update() bool {
	w, h := ebiten.WindowSize()
	if a.Health <= 0 || a.PositionX < 0-50 || a.PositionY < 0-50 || a.PositionX > float64(w)+50 || a.PositionY > float64(h)+50 {
		return false
	}
	if a.Health <= 60 {
		a.Sprite = assets.SpriteList.Asteroid[1]
	}
	if a.Health <= 30 {
		a.Sprite = assets.SpriteList.Asteroid[2]
	}
	a.PositionX += float64(a.VelocityX) * 3 / a.Mass // Velocity scales opposite of mass
	a.PositionY += float64(a.VelocityY) * 3 / a.Mass
	a.Rotation += a.VelocityR * 3 / int(a.Mass)
	return true
}
