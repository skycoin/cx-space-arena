package game

import (
	"fmt"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(scr *ebiten.Image) {
	ship1, _, err := ebitenutil.NewImageFromFile("assets/ship1.png")
	if err != nil {
		log.Fatal(err)
	}
	pl1 := g.World.Players[0]
	w, h := ship1.Size()
	optionsPlayer := ebiten.DrawImageOptions{}
	optionsPlayer.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	optionsPlayer.GeoM.Rotate((pl1.Rotation*360 + 90) * (math.Pi / 180))
	optionsPlayer.GeoM.Scale(0.1, 0.1)
	// TODO: Automatic color for each player
	// options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
	optionsPlayer.GeoM.Translate(pl1.PositionX, pl1.PositionY)
	scr.DrawImage(ship1, &optionsPlayer)
	ebitenutil.DebugPrint(scr, fmt.Sprint(ebiten.CurrentTPS()))
}
