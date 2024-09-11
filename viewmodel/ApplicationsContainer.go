package viewmodel

import (
	"keyboard-cheatsheet/main/ui"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewApplications(applications []string) *fyne.Container {
	hBox := container.NewHBox()
	for _, application := range applications {
		hBox.Add(NewApplication(application))
	}
	return hBox
}

func NewApplication(application string) fyne.CanvasObject {
	imagePath := "./icons/" + application + ".png"
	if FileExists(imagePath) {
		image := canvas.NewImageFromFile(imagePath)
		image.FillMode = canvas.ImageFillStretch
		image.SetMinSize(fyne.NewSize(30, 30))
		return image
	}
	return ui.NewText(application)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
