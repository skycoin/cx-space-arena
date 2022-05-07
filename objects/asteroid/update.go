package asteroid

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/assets"
)

func (asteroid *Asteroid) Update() bool {
	w, h := ebiten.WindowSize()
	if asteroid.Health <= 0 || asteroid.PositionX < 0-50 || asteroid.PositionY < 0-50 || asteroid.PositionX > float64(w)+50 || asteroid.PositionY > float64(h)+50 {
		return false
	}
	if asteroid.Health <= 60 {
		asteroid.Sprite = assets.SpriteList.Asteroid[1]
	}
	if asteroid.Health <= 30 {
		asteroid.Sprite = assets.SpriteList.Asteroid[2]
	}
	asteroid.PositionX += float64(asteroid.VelocityX) * 3 / asteroid.Mass // Velocity scales opposite of mass
	asteroid.PositionY += float64(asteroid.VelocityY) * 3 / asteroid.Mass
	asteroid.Rotation += asteroid.VelocityR * 3 / int(asteroid.Mass)
	return true
}
