package game

import (
	"fmt"
	"image/color"
	_ "image/png"
	"math"

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
			ebitenutil.DrawLine(scr, b.PositionX+math.Cos(float64(b.FireAngle)*(math.Pi/180))*25, b.PositionY+math.Sin(float64(b.FireAngle)*(math.Pi/180))*25, b.PositionX+math.Cos(float64(b.FireAngle)*(math.Pi/180))*500, b.PositionY+math.Sin(float64(b.FireAngle)*(math.Pi/180))*500, color.White)
		}
	}
	for _, v := range g.World.Asteroids {
		w, h := ebiten.WindowSize()
		if v != nil && v.PositionX > 0-25*v.Mass && v.PositionY > 0-25*v.Mass && v.PositionX < float64(w)+25*v.Mass && v.PositionY < float64(h)+25*v.Mass {
			v.Draw(scr)
		}
	}
	// Get FPS counter
	ebitenutil.DebugPrint(scr, fmt.Sprint(ebiten.CurrentTPS()))
}
