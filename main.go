package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/game"
	"github.com/skycoin/cx-space-arena/world"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano()) // Randomize the seed for asteroids, etc..
}

func main() {
	ebiten.SetWindowTitle("cx-space-arena")
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizable(true)

	g := game.Game{World: world.NewWorld()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
