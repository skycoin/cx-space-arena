package player

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	PlayerID             int
	PositionX, PositionY float64
	Velocity             int
	Rotation             int
	Health               int
	Points               int
	Deaths               int
	Cooldown             int
}

// Player contructor
func NewPlayer(ID int) *Player {
	w, h := ebiten.WindowSize()
	return &Player{
		PlayerID:  ID,
		PositionX: float64(w) / 2, PositionY: float64(h) / 2,
		Velocity: 0,
		Rotation: 0,
		Health:   100,
		Points:   0,
		Deaths:   0,
		Cooldown: 0,
	}
}
