package view

import (
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/linq"
	"keyboard-cheatsheet/main/ui"

	"fyne.io/fyne/v2"
)

func ToKeyCombinationViews(keyCombinations []data.KeyCombination) []KeyCombinationView {
	return linq.Map(keyCombinations, func(keyCombination data.KeyCombination) KeyCombinationView {
		return ToKeyCombinationView(keyCombination)
	})
}
func NewKeyCodeView(key string, isPressed bool) KeyCodeView {
	canvasText := ui.NewText(key + " ")
	canvasText.Alignment = fyne.TextAlignCenter
	return KeyCodeView{
		Key:        key,
		isPressed:  isPressed,
		canvasText: canvasText,
	}
}
func ToKeyCombinationView(keyCombination data.KeyCombination) KeyCombinationView {

	keys := linq.Map(keyCombination.Keys, func(key data.KeyCode) KeyCodeView {
		canvasText := ui.NewText(string(key) + " ")
		canvasText.Alignment = fyne.TextAlignCenter

		return KeyCodeView{
			Key:        string(key),
			isPressed:  false,
			canvasText: canvasText,
		}
	})
	keys = SortByPressedKeysCode(keys)

	keysText := ui.NewText(ToText(keys))

	return KeyCombinationView{
		Keys:                  keys,
		Description:           keyCombination.Description,
		Applications:          keyCombination.Applications,
		DescriptionText:       ui.NewText(keyCombination.Description),
		ApplicationsContainer: NewApplicationsView(keyCombination.Applications),
		KeysText:              keysText,
	}
}
