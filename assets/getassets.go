package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed particle.png
var particleImageBytes []byte

//go:embed particle2.png
var particle2ImageBytes []byte

//go:embed particle3.png
var particle3ImageBytes []byte

//go:embed particle4.png
var particle4ImageBytes []byte

// ParticleImage is a global variable to store the particle image
var ParticleImage *ebiten.Image

// Get loads the particle image into memory based on the configuration.
func Get() {
	var imageBytes []byte

	switch config.General.ParticleImage {
	case "particle.png":
		imageBytes = particleImageBytes
	case "particle2.png":
		imageBytes = particle2ImageBytes
	case "particle3.png":
		imageBytes = particle3ImageBytes
	case "particle4.png":
		imageBytes = particle4ImageBytes
	default:
		log.Fatal("Unknown image name: ", config.General.ParticleImage)
	}

	// Decode the embedded image
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		log.Fatal("Problem while decoding embedded particle image: ", err)
	}

	// Convert the image to an ebiten.Image
	ParticleImage = ebiten.NewImageFromImage(img)
}
