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
		if b != nil && b.Lifespan > 0 {
			b.Draw(scr)
		}
	}
	for _, a := range g.World.Asteroids {
		w, h := ebiten.WindowSize()
		if a != nil &&
			a.PositionX > 0-50 && // Offset unload by 50 to compensate for asteroid size
			a.PositionY > 0-50 &&
			a.PositionX < float64(w)+50 &&
			a.PositionY < float64(h)+50 {
			a.Draw(scr)
		}
	}
	// Get FPS counter
	ebitenutil.DebugPrint(scr, fmt.Sprint(ebiten.CurrentTPS()))
}
