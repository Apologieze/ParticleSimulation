package config

import "github.com/hajimehoshi/ebiten/v2"

//Fonction qui change de fichier config parmis une liste définit et avec un sens de défilement donné (shift positif ou négatif)
func ChangeConfig(shift int) {
	var configPreset []string = []string{"config.json", "config/config2.json", "config/config3.json", "config/config4.json", "config/config5.json"}
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
}
