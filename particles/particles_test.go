package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"testing"
)

// Test du nombre de particules à l'initialisation du système avec le fichier config
func TestInitSystem(t *testing.T) {
	config.Get("../config.json")
	var s System = NewSystem()
	if s.Content.Len() != config.General.InitNumParticles {
		t.Error("Votre fonction genere", s.Content.Len(), "particules au lieu de:", config.General.InitNumParticles)
	}
}

// Test d'un nombre aléatoire d'ajout de particule au système un grand nombre de fois
func TestAddingParticleRandom(t *testing.T) {
	var s System
	var randomTemp int
	for i := 0; i < 50; i++ {
		s = System{Content: list.New(), SpawnRate: 0, DeadParticles: list.New()}
		randomTemp = rand.Intn(500)
		s.Add(randomTemp)
		if s.Content.Len() != randomTemp {
			t.Error("Il devrait y avoir", randomTemp, "particules, mais il y en a:", s.Content.Len())
		}
	}
}

// Test du nombre de particule après 2 tick avec un SpawnRate entier aléatoire un grand nombre de fois
func TestSpawnRate1(t *testing.T) {
	var s System
	var randomTemp int
	for i := 0; i < 50; i++ {
		randomTemp = rand.Intn(50)
		s = System{Content: list.New(), SpawnRate: float64(randomTemp), DeadParticles: list.New()}
		s.SpawnUpdate()
		s.SpawnUpdate()
		if s.Content.Len() != randomTemp*2 {
			t.Error("Après 2 ticks d'update il devrait y avoir", randomTemp*2, "particules, mais nous en avons:", s.Content.Len())
		}
	}
}

// Test du nombre de particule après 100 ticks avec un SpawnRate d'un float
func TestSpawnRate2(t *testing.T) {
	var s System = System{Content: list.New(), SpawnRate: 0.11, DeadParticles: list.New()}
	for i := 0; i < 100; i++ {
		s.SpawnUpdate()
	}
	if s.Content.Len() != 11 {
		t.Error("Avec un SpawnRate de 0.11 au bout de 100 ticks on devrait avoir 11 particules, mais on en a:", s.Content.Len())
	}
}

// Test du mouvement d'une particule dans un sytème avec la méthode Move appelé depuis la méthode Update
func TestMovementSystem(t *testing.T) {
	var s System = System{Content: list.New(), SpawnRate: 0, DeadParticles: list.New(), RunningSimulation: true}
	var tempParticle Particle = Particle{
		PositionX: 0, PositionY: 1.1, SpeedX: 1.1, SpeedY: -1.1, Alive: true,
	}
	s.Content.PushFront(&tempParticle)
	s.Update()
	var px, py float64 = s.Content.Front().Value.(*Particle).PositionX, s.Content.Front().Value.(*Particle).PositionY
	if px != 1.1 || py != 0 {
		t.Error("La particule devrait se trouver en (1.1; 0) mais est en", "(", px, ";", py, ")")
	}
}

// Test sur le changement d'opacité d'une particule après une Update
func TestChangeOpacity(t *testing.T) {
	var tempParticle Particle
	var randomFloat float64
	for i := 0; i < 50; i++ {
		randomFloat = rand.Float64()
		tempParticle = Particle{Alive: true, Opacity: 0}
		tempParticle.ChangeOpacity(randomFloat)
		if tempParticle.Opacity != randomFloat {
			t.Error("Après une update la particule devrait avoir une opacité de", randomFloat, "mais nous en avons une opacité de", tempParticle.Opacity)
		}
	}
}

// Test sur le changement d'opacité d'une particule vérifiant qu'elle est bien "tuée" si son opacité est en dessous de 0.
func TestKillOpacity(t *testing.T) {
	var tempParticle Particle
	for i := 0; i < 50; i++ {
		tempParticle = Particle{Alive: true, Opacity: 0}
		var randfloat float64 = -(rand.Float64() + 0.01)
		tempParticle.ChangeOpacity(randfloat)
		if tempParticle.Alive == true {
			t.Error("L'opacité de la particule est de:", tempParticle.Opacity, "mais n'est pas tuée")
		}
	}
}

// Test qui vérifie que la gravité est bien appliqué sur la vitesse verticale de la particule
// (valeur de gravité calculé avec la fonction randomSpeed qui donne des valeurs positives et négatives)
func TestGravity(t *testing.T) {
	var tempParticle Particle = Particle{SpeedY: 0, Alive: true}
	for i := 0; i < 100; i++ {
		tempParticle.SpeedY = 0
		var randfloat float64 = randomSpeed()
		tempParticle.ApplyGravity(randfloat)
		if tempParticle.SpeedY != randfloat {
			t.Error("La vitesse speedY devrait être à:", randfloat, "mais est à:", tempParticle.SpeedY)
		}
	}
}

// Test qui vérifie que la particule est bien détecté lorsque elle sort de la bordure de l'écran
func TestOutofScreen(t *testing.T) {
	config.Get("../config.json")
	InitMargin()
	var tempParticle Particle = Particle{PositionX: BorderRight + 1, PositionY: 0, Alive: true}
	if !tempParticle.OutOfScreen() {
		t.Error("La particule se trouve en x:", tempParticle.PositionX, "avec la bordure en: ", BorderRight, "elle aurait donc du mourir")
	}
}
