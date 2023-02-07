package particles

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	if s.InputEnable {
		s.KeyCheck()
	}
	if s.RunningSimulation {
		s.SpawnUpdate()
		particleList := s.Content
		for element := particleList.Front(); element != nil; element = element.Next() {
			particle := element.Value.(*Particle)
			if particle.Alive {
				if s.ToggleGravity {
					particle.ApplyGravity(s.GravityPower)
				}
				particle.Move()
				if particle.OutOfScreen() || particle.ChangeOpacity(s.EvolveOpacity) {
					s.DeadParticles.PushFront(particle)
				}
				particle.ChangeColor()
			}
		}
	}
}
