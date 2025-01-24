package config

import "github.com/hajimehoshi/ebiten/v2"

var (
	config1Content = Config{
		WindowTitle:      "Project particles",
		WindowSizeX:      1900,
		WindowSizeY:      1000,
		ParticleImage:    "particle3.png",
		Debug:            true,
		InitNumParticles: 1,
		RandomSpawn:      false,
		SpawnX:           950,
		SpawnY:           500,
		SpawnRate:        30,
		RandomColor:      1,
		ColorNumber:      4,
		MarginBorder:     20,
		InputEnable:      true,
		GravityPower:     0.4,
		SpawnOpacity:     1,
		EvolveOpacity:    -0.007,
		ChangeColor:      0,
		StarMode:         false,
	}

	config2Content = Config{
		WindowTitle:      "Rain",
		WindowSizeX:      1900,
		WindowSizeY:      1000,
		ParticleImage:    "particle2.png",
		InitNumParticles: 1,
		RandomSpawn:      true,
		SpawnX:           950,
		SpawnY:           500,
		SpawnRate:        20,
		RandomColor:      0,
		ColorNumber:      2,
		MarginBorder:     20,
		InputEnable:      true,
		GravityPower:     0.6,
		SpawnOpacity:     1,
		EvolveOpacity:    -0.03,
		ChangeColor:      0,
		StarMode:         false,
	}

	config3Content = Config{
		WindowTitle:      "Fire",
		WindowSizeX:      1900,
		WindowSizeY:      1000,
		ParticleImage:    "particle.png",
		InitNumParticles: 1,
		RandomSpawn:      false,
		SpawnX:           950,
		SpawnY:           800,
		SpawnRate:        100,
		RandomColor:      2,
		ColorNumber:      1,
		MarginBorder:     20,
		InputEnable:      true,
		GravityPower:     -0.5,
		SpawnOpacity:     1,
		EvolveOpacity:    -0.02,
		ChangeColor:      -0.017,
		StarMode:         false,
	}

	config4Content = Config{
		WindowTitle:      "Stars",
		WindowSizeX:      1900,
		WindowSizeY:      1000,
		ParticleImage:    "particle4.png",
		InitNumParticles: 1,
		RandomSpawn:      true,
		SpawnX:           950,
		SpawnY:           500,
		SpawnRate:        0.5,
		RandomColor:      0,
		ColorNumber:      4,
		MarginBorder:     20,
		InputEnable:      true,
		GravityPower:     0,
		SpawnOpacity:     1,
		EvolveOpacity:    -0.005,
		ChangeColor:      0,
		StarMode:         true,
	}

	config5Content = Config{
		WindowTitle:      "Christmas Lights",
		WindowSizeX:      1900,
		WindowSizeY:      1000,
		ParticleImage:    "particle.png",
		InitNumParticles: 1,
		RandomSpawn:      true,
		SpawnX:           950,
		SpawnY:           500,
		SpawnRate:        100,
		RandomColor:      1,
		ColorNumber:      1,
		MarginBorder:     20,
		InputEnable:      true,
		GravityPower:     0.5,
		SpawnOpacity:     1,
		EvolveOpacity:    -0.1,
		ChangeColor:      0,
		StarMode:         false,
	}

	configArray = []Config{config2Content, config3Content, config4Content, config1Content, config5Content}
)

// Fonction qui change de fichier config parmis une liste définit et avec un sens de défilement donné (shift positif ou négatif)
/*func ChangeConfig(shift int) {
	var configPreset []string = []string{"config/config2.json", "config/config3.json", "config/config4.json", "config1.json", "config/config5.json"}
	presetNumber += shift
	if presetNumber == len(configPreset) {
		presetNumber = 0
	}
	if presetNumber == -1 {
		presetNumber = len(configPreset) - 1
	}
	Get(configPreset[presetNumber])
	ebiten.SetWindowTitle(General.WindowTitle)
	ebiten.SetWindowSize(General.WindowSizeX, General.WindowSizeY)
}*/

func ChangeConfig(shift int) {
	presetNumber += shift
	if presetNumber == len(configArray) {
		presetNumber = 0
	}
	if presetNumber == -1 {
		presetNumber = len(configArray) - 1
	}
	General = configArray[presetNumber]
	ebiten.SetWindowTitle(General.WindowTitle)
	ebiten.SetWindowSize(General.WindowSizeX, General.WindowSizeY)
}
