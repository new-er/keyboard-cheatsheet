package view

import (
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/linq"
	"keyboard-cheatsheet/main/ui"
)

func ToKeyCombinationViews(keyCombinations []data.KeyCombination, pressedKeys []data.KeyCode) []KeyCombinationView {
	return linq.Map(keyCombinations, func(keyCombination data.KeyCombination) KeyCombinationView {
		return ToKeyCombinationView(keyCombination, pressedKeys)
	})
}

func ToKeyCombinationView(keyCombination data.KeyCombination, pressedKeys []data.KeyCode) KeyCombinationView {
	keys := linq.Map(keyCombination.Keys, func(key data.KeyCode) KeyCodeView {
		return KeyCodeView{
			Key:       string(key),
			IsPressed: Contains(pressedKeys, key),
		}
	})
	keys = SortByPressedKeysCode(keys)

	return KeyCombinationView{
		Keys:                  keys,
		Description:           keyCombination.Description,
		Applications:          keyCombination.Applications,
		DescriptionText:       ui.NewText(keyCombination.Description),
		ApplicationsContainer: NewApplicationsView(keyCombination.Applications),
	}
}
