package particles

import (
	"math"
)

// Méthode qui vérifie à chaque tick si de nouvelles particules doivent apparaitre et actualise la valeur de Spawn
func (s *System) SpawnUpdate() {
	s.SpawnValue += s.SpawnRate
	if s.SpawnValue >= 1 {
		if s.InputEnable {
			s.MoveGenerator()
		}
		s.Add(int(s.SpawnValue))
		_, s.SpawnValue = math.Modf(s.SpawnValue)
	}
}

// Méthode appelé sur une particule "morte" pour la refaire apparaitre en tant que nouvelle particule et optimiser la mémoire
func (p *Particle) ReSpawn(posx, posy, oppacity float64, rgb Color) {
	p.Alive = true
	p.PositionX, p.PositionY, p.Opacity = posx, posy, oppacity
	p.ColorRed, p.ColorGreen, p.ColorBlue = rgb.red, rgb.green, rgb.blue
	p.SpeedX, p.SpeedY = randomSpeed(), randomSpeed()
}
