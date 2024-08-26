package main

import (
	"fmt"
)

type ApplicationState struct {
	ActiveWindowTitle          string
	AllKeyCombinations         []KeyCombination
	TransformedKeyCombinations []IsPressedKeyCombination
	PressedKeys                []KeyCode
	Error                      string
}

var (
	renderToggle = false
)

func getRenderToggleText() string {
	if renderToggle {
		return "•"
	}
	return "◦"
}

func ConsoleWriteApplicationState(state ApplicationState) {
	ClearLines()
	errorText := ""
	if len(state.Error) > 0 {
		errorText = fmt.Sprint(colorRed, "[", state.Error, "]", colorReset)
	}
	WriteLine(fmt.Sprintf("[%s] Keyboard Cheatsheet %s", getRenderToggleText(), errorText))

	WriteLine(state.ActiveWindowTitle)
	WriteLine(fmt.Sprint(state.PressedKeys))

	for _, keyCombination := range state.TransformedKeyCombinations {
		ConsoleWriteKeyCombination(keyCombination)
	}
}

var (
	colorReset = "\u001b[0m"
	colorRed   = "\u001b[31m"
	colorGreen = "\u001b[32m"
)

func ConsoleWriteKeyCombination(keyCombination IsPressedKeyCombination) {
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
		if key == item {
			return true
		}
	}

	return false
}
