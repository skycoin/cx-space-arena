package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/consts"
)

// Player draw method
func (player *Player) Draw(screen *ebiten.Image) {
	spriteWidth, spriteHeight := player.Sprite.Size()
	options := ebiten.DrawImageOptions{} // Translate image to rotate along center
	options.GeoM.Translate(-float64(spriteWidth)/2, -float64(spriteHeight)/2)
	options.GeoM.Rotate(float64(player.Rotation+90) * (math.Pi / 180))
	options.GeoM.Scale(consts.PLAYER_SCALE, consts.PLAYER_SCALE)
	options.GeoM.Translate(player.PositionX, player.PositionY)
	screen.DrawImage(player.Sprite, &options)
}
