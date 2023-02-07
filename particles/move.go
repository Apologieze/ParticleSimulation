package particles

// Méthode qui actualise la position de la particule par rapport à sa vitesse
func (p *Particle) Move() {
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
}

// Méthode qui applique la force de gravité à la vitesse de la particule
// Est appelé seulement si la gravité est activé (par le clavier ou par défaut)
func (p *Particle) ApplyGravity(gravity float64) {
	p.SpeedY += gravity
}
