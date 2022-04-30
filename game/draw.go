package game

import (
	"fmt"
	"image/color"
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
	// TESTING: Drawing hitboxes, remove later!
	for _, b := range g.World.Bullets {
		if b != nil {
			b.Draw(scr)
			ebitenutil.DrawRect(scr, b.PositionX-25, b.PositionY-25, 50, 50, color.RGBA{0, 255, 0, 100})
		}
	}
	for _, a := range g.World.Asteroids {
		if a != nil {
			a.Draw(scr)
			ebitenutil.DrawRect(scr, a.PositionX-12.5*a.Mass, a.PositionY-12.5*a.Mass, 25*a.Mass, 25*a.Mass, color.RGBA{255, 0, 0, 100})
		}
	}
	// Get FPS counter
	ebitenutil.DebugPrint(scr, fmt.Sprint(ebiten.CurrentTPS()))

}
