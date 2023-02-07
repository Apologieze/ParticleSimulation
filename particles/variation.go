package particles

import (
	"project-particles/config"
)

// Méthode faisant varié l'Opacité de la particule en fonction de EvoleOpacity
func (p *Particle) ChangeOpacity(value float64) (killed bool) {
	if value != 0 {
		var total float64 = p.Opacity + value
		if total >= 0 && total <= 1 {
			p.Opacity = total
		} else if total < 0 {
			p.Kill()
			killed = true
		}
	}
	return killed
}

// Permet de modifier la couleur d'une particule à chaque tick
func (p *Particle) ChangeColor() {
	switch config.General.RandomColor {
	case 0:
		return
	case 2:
		if p.ColorGreen != 0 && p.ColorGreen != 1 {
			p.ColorGreen += config.General.ChangeColor
			if p.ColorGreen > 1 {
				p.ColorGreen = 1
			} else if p.ColorGreen < 0 {
				p.ColorGreen = 0
			}
		}
	}
}
