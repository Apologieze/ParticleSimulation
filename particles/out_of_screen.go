package particles

import (
	"project-particles/config"
)

var BorderLeft, BorderRight, BorderTop, BorderBottom float64

// Initialisé une fois dans le New et définie la bordure dans laquelle les particules sont en vie
func InitMargin() {
	BorderLeft = float64(-config.General.MarginBorder)
	BorderRight = float64(config.General.WindowSizeX + config.General.MarginBorder)
	BorderBottom = float64(config.General.WindowSizeY + config.General.MarginBorder)
	BorderTop = float64(-config.General.MarginBorder)
}

// Méthode vérifiant si la particule est toujours dans l'écran et sinon tue l'écran
func (p *Particle) OutOfScreen() (killed bool) {
	if p.PositionX < BorderLeft || p.PositionX > BorderRight || p.PositionY < BorderTop || p.PositionY > BorderBottom {
		p.Kill()
		killed = true
	}
	return killed
}
