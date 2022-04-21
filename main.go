package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/game"
	"github.com/skycoin/cx-space-arena/world"
)

func main() {
	ebiten.SetWindowTitle("cx-space-arena")
	ebiten.SetWindowSize(640, 480)

	g := game.Game{World: world.NewWorld()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
