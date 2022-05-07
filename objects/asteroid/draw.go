package asteroid

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/consts"
)

// Asteroid draw method
func (asteroid *Asteroid) Draw(screen *ebiten.Image) {
	spriteWidth, spriteHeight := asteroid.Sprite.Size()
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(spriteWidth)/2, -float64(spriteHeight)/2)
	options.GeoM.Rotate(float64(asteroid.Rotation) * (math.Pi / 180))
	options.GeoM.Scale(consts.ASTEROID_SCALE*asteroid.Mass, consts.ASTEROID_SCALE*asteroid.Mass) // Size between 0.1 and 0.2
	options.GeoM.Translate(asteroid.PositionX, asteroid.PositionY)
	screen.DrawImage(asteroid.Sprite, &options)
}
