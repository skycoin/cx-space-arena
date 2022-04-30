package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/game"
	"github.com/skycoin/cx-space-arena/world"
)

func main() {
	ebiten.SetWindowTitle("cx-space-arena")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizable(true)

	rand.Seed(time.Now().UTC().UnixNano()) // Randomize the seed for asteroids, etc..

	g := game.Game{World: world.NewWorld()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
