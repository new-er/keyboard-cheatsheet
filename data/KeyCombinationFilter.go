package data

import "keyboard-cheatsheet/main/linq"

func FilterDisabledKeyCombinations(k []KeyCombination) []KeyCombination {
	return linq.Filter(k, func(keyCombination KeyCombination) bool {
		return !keyCombination.Disabled
	})
}
