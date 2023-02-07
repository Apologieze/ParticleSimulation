package particles

import (
	"math/rand"
	"project-particles/config"
)

type Color struct {
	red   float64
	green float64
	blue  float64
}

// Choisie quelle fonction definira la couleur entre des preset de couleur fixe ou des fonctions de couleur aléatoire
// (Pour choisir cela prend en valeur la variable RandomColor de la config)
func chooseColor() Color {
	switch config.General.RandomColor {
	case 0:
		return choosePreset()
	case 1:
		return randomColorRainbow()
	case 2:
		return randomColorFire()
	default:
		return choosePreset()
	}
}

// Choisie une couleur fixe parmis des preset de couleur
func choosePreset() Color {

	var colorPreset []Color = []Color{{1, 1, 1}, {1, 0, 0}, {0, 0.58, 1}, {1, 0.25, 0.77}, {1, 1, 0}}
	var indexColor int = config.General.ColorNumber
	if indexColor >= len(colorPreset) {
		indexColor = 0
	}
	return colorPreset[indexColor]
}

// Retourne une couleur avec une teinte aléatoire mais qui a forcément une saturation et une luminosité maximum (cela évite les différents niveaux de noir)
func randomColorRainbow() Color {
	var randCase int = rand.Intn(6)
	var randFloat float64 = rand.Float64()

	switch randCase {
	case 0:
		return Color{1, randFloat, 0}
	case 1:
		return Color{1, 0, randFloat}
	case 2:
		return Color{randFloat, 1, 0}
	case 3:
		return Color{0, 1, randFloat}
	case 4:
		return Color{randFloat, 0, 1}
	case 5:
		return Color{0, randFloat, 1}
	}
	return Color{1, 1, 1}
}

// Retourne une couleur entre le orange et le jaune
// La valeur de vert entre 0.7 et 1 est arbitraire et a été définie après des tests
func randomColorFire() Color {
	var randFloat float64 = 0.7 + rand.Float64()*(0.3)
	if randFloat > 1 {
		randFloat = 1
	}
	return Color{1, randFloat, 0}
}
