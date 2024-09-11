package ui

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func NewText(text string) *canvas.Text {
	canvasText := canvas.NewText(text, color.White)
	canvasText.TextSize = 30
	return canvasText
}
