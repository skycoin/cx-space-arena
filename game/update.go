package game

// Update function, called 60 times each second
func (g *Game) Update() error {
	g.World.CurrentTick++
	if g.World.CurrentTick%60 == 0 {
		g.World.CurrentTick = 0
	}
	g.UpdateInputs()
	for _, player := range g.World.Players {
		if player != nil {
			player.Update()
		}
	}
	for i, bullet := range g.World.Bullets {
		if bullet != nil {
			if !bullet.Update() {
				g.World.Bullets[i] = nil
			}
		}
	}
	for i, asteroid := range g.World.Asteroids {
		if asteroid != nil {
			if !asteroid.Update() {
				g.World.Asteroids[i] = nil
			}
		}
	}
	g.UpdateCollisions()
	return nil
}
