package bullet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/constants"
)

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	w, h := bullet.Sprite.Size()
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	options.GeoM.Rotate(float64(bullet.Angle) * (math.Pi / 180))
	options.GeoM.Scale(constants.BULLET_SCALE, constants.BULLET_SCALE)
	options.GeoM.Translate(bullet.PositionX, bullet.PositionY)
	screen.DrawImage(bullet.Sprite, &options)
}
