package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(scr *ebiten.Image) {
	for i := 0; i < len(g.World.Players); i++ {
		scr.DrawImage()
	}
	ebitenutil.DebugPrint(scr, fmt.Sprint(ebiten.CurrentTPS()))
}
