package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			if p.Alive {
				options := ebiten.DrawImageOptions{}
				options.GeoM.Rotate(p.Rotation)
				options.GeoM.Scale(p.ScaleX, p.ScaleY)
				options.GeoM.Translate(p.PositionX, p.PositionY)
				options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
				screen.DrawImage(assets.ParticleImage, &options)
			}
		}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint("FPS: ", ebiten.ActualTPS()))
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Particles: ", g.system.Content.Len()), 0, 13)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Dead Particles: ", g.system.DeadParticles.Len()), 0, 26)
		ebitenutil.DebugPrintAt(screen, "Touche [S] pour afficher/cacher les stats", 820, 0)
		ebitenutil.DebugPrintAt(screen, "Touche [G] pour activer la gravité", 820, 13)
		ebitenutil.DebugPrintAt(screen, "Touche [P] pour mettre en pause", 820, 26)
		ebitenutil.DebugPrintAt(screen, "Touche [Espace] pour clear", 820, 40)
		ebitenutil.DebugPrintAt(screen, "Touche [flèche gauche/doite] pour changer de config", 820, 53)
	}

}
