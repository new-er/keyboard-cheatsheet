package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type KeyCombinationView struct {
	Keys         []KeyCodeView
	Description  string
	Applications []string

	ApplicationsContainer *fyne.Container
	DescriptionText       *canvas.Text
}
