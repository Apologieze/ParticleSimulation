package particles

import (
	"math/rand"
	"project-particles/config"
)

// créer une vitesse aléatoire entre -4 et 4 (deux chiffre prient par expérimentation pour rendre cela intéressant visuellement)
// avec une valeur minimum pour ne pas avoir de particules immobile
// Appelle la fonction StarSpeed si le mode étoile est activé
func randomSpeed() float64 {
	if config.General.StarMode {
		return StarSpeed()
	}
	var value float64 = float64(rand.Intn(3)) + rand.Float64() + 0.01
	if rand.Intn(2) == 0 {
		return value
	}
	return -value
}

// Méthode qui ajoute un nombre de particule "number" au System fait pour être le plus optimisé possible
// et limiter le nombre d'appel par tick à cette méthode
func (s *System) Add(number int) {
	var posx, posy float64 = float64(s.PositionX), float64(s.PositionY)

	for i := 0; i < number; i++ {
		if config.General.RandomSpawn {
			if config.General.StarMode {
				posx, posy = StarPos()
			} else {
				posx = float64(rand.Intn(config.General.WindowSizeX))
				posy = float64(rand.Intn(config.General.WindowSizeY))
			}
		}

		var tempColor Color = chooseColor()
		if s.DeadParticles.Len() == 0 {

			s.Content.PushFront(&Particle{
				PositionX: posx,
				PositionY: posy,
				ScaleX:    1, ScaleY: 1,
				ColorRed: tempColor.red, ColorGreen: tempColor.green, ColorBlue: tempColor.blue,
				Opacity: s.SpawnOpacity,
				SpeedX:  randomSpeed(),
				SpeedY:  randomSpeed(),
				Alive:   true,
			})
		} else {
			tempParticle := s.DeadParticles.Front()
			tempParticle.Value.(*Particle).ReSpawn(posx, posy, s.SpawnOpacity, tempColor)
			s.DeadParticles.Remove(tempParticle)
		}
	}
}
