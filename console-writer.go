package main

import (
	"fmt"
)

type ApplicationState struct {
	ActiveWindowTitle          string
	AllKeyCombinations         []KeyCombination
	TransformedKeyCombinations []KeyCombinationView
	PressedKeys                []KeyCode
	Error                      string
}

var (
	renderToggle = false
)

func getRenderToggleText() string {
	text := "◦"
	if renderToggle {
		text = "•"
	}
	renderToggle = !renderToggle
	return text
}

func ConsoleWriteApplicationState(state ApplicationState) {
	ClearLines()
	errorText := ""
	if len(state.Error) > 0 {
		errorText = fmt.Sprint(colorRed, "[", state.Error, "]", colorReset)
	}
	WriteLine(fmt.Sprintf("[%s] Keyboard Cheatsheet %s %s [%s]", getRenderToggleText(), errorText, fmt.Sprint(state.PressedKeys), state.ActiveWindowTitle))

	for _, keyCombination := range state.TransformedKeyCombinations {
		ConsoleWriteKeyCombination(keyCombination)
	}
}

var (
	colorReset = "\u001b[0m"
	colorRed   = "\u001b[31m"
	colorGreen = "\u001b[32m"
)

func ConsoleWriteKeyCombination(keyCombination KeyCombinationView) {
	keyCombinationDescription := keyCombination.description
	keyCombinationApplication := keyCombination.application
	keyCombinationKeys := []string{}
	for _, key := range keyCombination.keys {
		color := colorReset
		if key.isPressed {
			color = colorGreen
		}

		keyCombinationKeys = append(keyCombinationKeys, fmt.Sprint(color, key.key, colorReset))
	}

	text := fmt.Sprintf("%s: %s (%s)", keyCombinationKeys, keyCombinationDescription, keyCombinationApplication)
	WriteLine(text)
}

func Contains(slice []KeyCode, item KeyCode) bool {
	for _, key := range slice {
		if item.Matches(key) {
			return true
		}
	}

	return false
}
