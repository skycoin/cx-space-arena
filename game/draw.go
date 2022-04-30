package game

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw function, called 60 times each second
func (g *Game) Draw(scr *ebiten.Image) {
	for _, p := range g.World.Players {
		if p != nil {
			p.Draw(scr)
		}
	}
	for _, b := range g.World.Bullets {
		if b != nil {
			b.Draw(scr)
		}
	}
	for _, a := range g.World.Asteroids {
		if a != nil {
			a.Draw(scr)
		}
	}
	// Get FPS counter
	ebitenutil.DebugPrint(scr, fmt.Sprint(ebiten.CurrentTPS()))
}
