package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var (
	icons = make(map[string]fyne.CanvasObject)
)

func GetImageOrText(application string) fyne.CanvasObject {
	// Check if the icon is already cached
	if icon, exists := icons[application]; exists {
		return icon
	}

	imagePath := "./icons/" + application + ".png"
	if FileExists(imagePath) {
		image := canvas.NewImageFromFile(imagePath)
		image.FillMode = canvas.ImageFillStretch
		image.SetMinSize(fyne.NewSize(30, 30))
		icons[application] = image // Cache the image
		return image
	}

	text := NewText(application)
	icons[application] = text // Cache the text
	return text
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
