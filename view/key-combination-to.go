package view

import (
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/linq"
	"keyboard-cheatsheet/main/ui"
)

func ToKeyCombinationViews(keyCombinations []data.KeyCombination) []KeyCombinationView {
	return linq.Map(keyCombinations, func(keyCombination data.KeyCombination) KeyCombinationView {
		return ToKeyCombinationView(keyCombination)
	})
}

func ToKeyCombinationView(keyCombination data.KeyCombination) KeyCombinationView {
	keys := linq.Map(keyCombination.Keys, func(key data.KeyCode) KeyCodeView {
		return KeyCodeView{
			Key:       string(key),
			IsPressed: false,
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
