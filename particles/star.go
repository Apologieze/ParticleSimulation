package particles

import (
	"math/rand"
	"project-particles/config"
)

// Donne une position de spawn pour les étoiles (sur le bord gauche ou haut de l'écran)
func StarPos() (x, y float64) {
	if rand.Intn(2) == 0 {
		return float64(rand.Intn(config.General.WindowSizeX)), 0
	} else {
		return 0, float64(rand.Intn(config.General.WindowSizeY))
	}
}

// Donne un vitesse aux étoiles (toujours vers le droite et le bas de l'écran)
func StarSpeed() float64 {
	return float64(rand.Intn(2)+10) + rand.Float64()
}
