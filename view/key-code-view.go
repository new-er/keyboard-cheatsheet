package view

import (
	"image/color"
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/linq"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type KeyCodeView struct {
	Key        string
	isPressed  bool
	canvasText *canvas.Text
}

func ToText(keyCodes []KeyCodeView) string {
	return strings.Join(linq.Map(keyCodes, func(key KeyCodeView) string {
		return string(key.Key)
	}), " + ")
}
func Contains(slice []data.KeyCode, item data.KeyCode) bool {
	for _, key := range slice {
		if item.MatchesKeyCode(key) {
			return true
		}
	}

	return false
}

func (k *KeyCodeView) IsPressed() bool {
	return k.isPressed
}

func (k *KeyCodeView) SetIsPressed(isPressed bool) {
	k.isPressed = isPressed

	if isPressed {
		k.canvasText.Color = color.RGBA{0, 255, 0, 255}
		k.canvasText.TextStyle = fyne.TextStyle{Bold: true}
	} else {
		k.canvasText.Color = color.White
		k.canvasText.TextStyle = fyne.TextStyle{Bold: false}
	}
}

func (k *KeyCodeView) CanvasText() *canvas.Text {
	return k.canvasText
}
