package view

import (
	"keyboard-cheatsheet/main/data"
)

func UpdatePressedKeys(keyCombinations []KeyCombinationView, pressedKeys []data.KeyCode) []KeyCombinationView {
	for i := range keyCombinations {
		keyCombinations[i].Keys = updatePressedKeyCodeViews(keyCombinations[i].Keys, pressedKeys)
	}

	return keyCombinations
}

func updatePressedKeyCodeViews(keys []KeyCodeView, pressedKeys []data.KeyCode) []KeyCodeView {
	for i := range keys {
		keys[i].IsPressed = isPressed(pressedKeys, keys[i])
	}

	return keys
}

func isPressed(slice []data.KeyCode, item KeyCodeView) bool {
	for _, key := range slice {
		if data.MatchesKeyString(string(key), item.Key) {
			return true
		}
	}

	return false
}
