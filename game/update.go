package game

// Update function, called 60 times each second
func (g *Game) Update() error {
	g.UpdateInputs()
	for _, p := range g.World.Players {
		if p != nil {
			p.Update()
		}
	}
	for i, b := range g.World.Bullets {
		if b != nil {
			if !b.Update() {
				g.World.Bullets[i] = nil
			}
		}
	}
	for i, a := range g.World.Asteroids {
		if a != nil {
			if !a.Update() {
				g.World.Asteroids[i] = nil
			}
		}
	}
	g.UpdateCollisions()
	return nil
}
