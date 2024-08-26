package main

import (
	"sort"
)

type KeyCodeView struct {
	key       KeyCode
	isPressed bool
}

type KeyCombinationView struct {
	keys        []KeyCodeView
	description string
	application string
}

func ToKeyCombinationViews(keyCombinations []KeyCombination, pressedKeys []KeyCode) []KeyCombinationView {
	return Map(keyCombinations, func(keyCombination KeyCombination) KeyCombinationView {
		return ToKeyCombinationView(keyCombination, pressedKeys)
	})
}

func ToKeyCombinationView(keyCombination KeyCombination, pressedKeys []KeyCode) KeyCombinationView {
	keys := Map(keyCombination.Keys, func(key KeyCode) KeyCodeView {
		return KeyCodeView{
			key:       key,
			isPressed: Contains(pressedKeys, key),
		}
	})

	return KeyCombinationView{
		keys:        keys,
		description: keyCombination.Description,
		application: keyCombination.Application,
	}
}

func SortByPressedKeys(views []KeyCombinationView) []KeyCombinationView {
	sort.Slice(views, func(i, j int) bool {
		countI := Count(views[i].keys, func(key KeyCodeView) bool {
			return key.isPressed
		})
		countJ := Count(views[j].keys, func(key KeyCodeView) bool {
			return key.isPressed
		})
		return countI > countJ
	})

	return views
}
