package game

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw function, called 60 times each second
func (g *Game) Draw(screen *ebiten.Image) {
	for _, player := range g.World.Players {
		if player != nil {
			if player.Recoil > 0 && g.World.CurrentTick&2 == 0 {
				continue
			}
			player.Draw(screen)
		}
	}
	for _, bullet := range g.World.Bullets {
		if bullet != nil {
			bullet.Draw(screen)
		}
	}
	for _, asteroid := range g.World.Asteroids {
		if asteroid != nil {
			asteroid.Draw(screen)
		}
	}
	// Get FPS counter
	ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))

}
