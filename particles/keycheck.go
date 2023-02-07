package particles

import (
	"container/list"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Méthode qui permet lors d'un click gauche de la souris (qui peut être maintenu) de déplacer le générateur de particule à l'emplacement du curseur
// (seulement lorsque InputEnable est activé dans la config)
func (s *System) MoveGenerator() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.PositionX, s.PositionY = ebiten.CursorPosition()
	}
}

// Méthode qui active ou désactive la gravité à chaque appuis sur la touche G du clavier
// Méthode qui active ou désactive l'affichage des statistiques à chaque appuis sur la touche S du clavier
// Méthode qui met la simulation en pause à chaque appuis sur la touche P du clavier
// Méthode qui permute de fichier config à chaque appuis sur les flèches droite/gauches
// Méthode qui clear toutes les particules vivantes et mortes à chaque appuis de la touche Espace
// (seulement lorsque InputEnable est activé dans la config)
func (s *System) KeyCheck() {
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		s.ToggleGravity = !s.ToggleGravity
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		config.General.Debug = !config.General.Debug
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		s.RunningSimulation = !s.RunningSimulation
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		config.ChangeConfig(1)
		s.systemConfigChange()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		config.ChangeConfig(-1)
		s.systemConfigChange()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.Content = list.New()
		s.DeadParticles = list.New()
	}
}
