package particles

import (
	"container/list"
	"math/rand"
	"project-particles/assets"
	"project-particles/config"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

func NewSystem() System {
	rand.Seed(time.Now().UnixNano())
	InitMargin()
	l := list.New()

	if config.General.SpawnOpacity == 0 {
		config.General.SpawnOpacity = 1
	}
	s := System{
		Content:   l,
		SpawnRate: config.General.SpawnRate,
		PositionX: config.General.SpawnX, PositionY: config.General.SpawnY,
		InputEnable:       config.General.InputEnable,
		ToggleGravity:     true,
		GravityPower:      config.General.GravityPower,
		SpawnOpacity:      config.General.SpawnOpacity,
		EvolveOpacity:     config.General.EvolveOpacity,
		DeadParticles:     list.New(),
		RunningSimulation: true,
	}
	s.Add(config.General.InitNumParticles)
	return s
}

// Reactualise le system après avoir changé de fichier de config
func (s *System) systemConfigChange() {
	s.SpawnRate = config.General.SpawnRate
	s.PositionX, s.PositionY = config.General.SpawnX, config.General.SpawnY
	s.InputEnable = config.General.InputEnable
	s.GravityPower = config.General.GravityPower
	s.SpawnOpacity = config.General.SpawnOpacity
	s.EvolveOpacity = config.General.EvolveOpacity
	s.ToggleGravity = true
	assets.Get()
}
