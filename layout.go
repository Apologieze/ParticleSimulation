package main

import "github.com/hajimehoshi/ebiten/v2"

// Layout définit la taille en pixels de la zone d'affichage des particules
// en fonction de la taille en pixels de la fenêtre.
func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}
