package particles

//Méthode qui modifie la particule la rendant morte
func (p *Particle) Kill() {
	p.Alive = false
}
