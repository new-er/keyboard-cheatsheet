package viewmodel

import (
	"image/color"
	"keyboard-cheatsheet/main/repository"
	"keyboard-cheatsheet/main/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type KeyCodeViewModel struct {
	text *canvas.Text
}

func NewKeyCodeViewModel(keyCode *repository.KeyCodeRepository) *KeyCodeViewModel {
	text := ui.NewText(string(keyCode.Key))
	listenToKeyCodeRepositoryChanges(keyCode, text)
	return &KeyCodeViewModel{
		text: text,
	}
}

func listenToKeyCodeRepositoryChanges(keyCode *repository.KeyCodeRepository, text *canvas.Text) {
	isPressed := keyCode.IsPressed.Observe()
	go func() {
		for {
			pressed := <-isPressed.Changes()
			if pressed {
				text.Color = color.RGBA{0, 255, 0, 255}
				text.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				text.Color = color.White
				text.TextStyle = fyne.TextStyle{Bold: false}
			}
		}
	}()

}
