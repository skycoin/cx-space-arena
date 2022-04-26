package game

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw function, called 60 times each second
func (g *Game) Draw(scr *ebiten.Image) {
	ship1, _, err := ebitenutil.NewImageFromFile("assets/ship1.png")
	if err != nil {
		log.Fatal(err)
	}
	asteroid, _, err := ebitenutil.NewImageFromFile("assets/asteroid.png")
	if err != nil {
		log.Fatal(err)
	}
	pl1 := g.World.Players[0]
	w, h := ship1.Size()
	optionsPlayer := ebiten.DrawImageOptions{}

	// Translate image to rotate along center
	optionsPlayer.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	optionsPlayer.GeoM.Rotate(float64(pl1.Rotation+90) * (math.Pi / 180))

	optionsPlayer.GeoM.Scale(0.1, 0.1)
	optionsPlayer.GeoM.Translate(pl1.PositionX, pl1.PositionY)
	// TODO: Automatic color for each player
	// options.ColorM.Scale()
	scr.DrawImage(ship1, &optionsPlayer)

	for _, v := range g.World.Bullets {
		if v != nil && v.Lifespan > 0 {
			ebitenutil.DrawLine(scr, v.PositionX+math.Cos(float64(v.FireAngle)*(math.Pi/180))*25, v.PositionY+math.Sin(float64(v.FireAngle)*(math.Pi/180))*25, v.PositionX+math.Cos(float64(v.FireAngle)*(math.Pi/180))*500, v.PositionY+math.Sin(float64(v.FireAngle)*(math.Pi/180))*500, color.White)
		}
	}
	w, h = ebiten.WindowSize()
	for _, v := range g.World.Asteroids {
		if v != nil && v.PositionX > 0-25*v.Mass && v.PositionY > 0-25*v.Mass && v.PositionX < float64(w)+25*v.Mass && v.PositionY < float64(h)+25*v.Mass {
			optionsAsteroid := ebiten.DrawImageOptions{}
			optionsAsteroid.GeoM.Translate(-float64(w)/2, -float64(h)/2)
			optionsAsteroid.GeoM.Rotate(float64(v.Rotation) * (math.Pi / 180))
			optionsAsteroid.GeoM.Scale(0.1*v.Mass, 0.1*v.Mass)
			optionsAsteroid.GeoM.Translate(v.PositionX, v.PositionY)
			scr.DrawImage(asteroid, &optionsAsteroid)
		}
	}

	// Get FPS counter
	ebitenutil.DebugPrint(scr, fmt.Sprint(ebiten.CurrentTPS()))
}
