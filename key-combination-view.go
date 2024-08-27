package main

import (
	"sort"
	"strings"
)

type KeyCodeView struct {
	key       string
	isPressed bool
}
type KeyCombinationView struct {
	keys        []KeyCodeView
	description string
	Application string
}

func ToKeyCombinationViews(keyCombinations []KeyCombination, pressedKeys []KeyCode) []KeyCombinationView {
	return Map(keyCombinations, func(keyCombination KeyCombination) KeyCombinationView {
		return ToKeyCombinationView(keyCombination, pressedKeys)
	})
}

func ToKeyCombinationView(keyCombination KeyCombination, pressedKeys []KeyCode) KeyCombinationView {
	keys := Map(keyCombination.Keys, func(key KeyCode) KeyCodeView {
		return KeyCodeView{
			key:       string(key),
			isPressed: Contains(pressedKeys, key),
		}
	})

	return KeyCombinationView{
		keys:        keys,
		description: keyCombination.Description,
		Application: keyCombination.Application,
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

func ToText(keyCodes []KeyCodeView) string {
	return strings.Join(Map(keyCodes, func(key KeyCodeView) string {
		return string(key.key)
	}), " + ")
}
func Contains(slice []KeyCode, item KeyCode) bool {
	for _, key := range slice {
		if item.Matches(key) {
			return true
		}
	}

	return false
}
