package main

import (
	"context"
	"time"
)

const (
	combinationsFile = "combinations.json"
)

func main() {
	combinations := KeyCombinationsFromFileOrPanic(combinationsFile)
	activeWindowChannel := GetActiveWindowTitleChannel()
	activeWindow := ""
	pressedKeysChannel := GetPressedKeysChannel()
	pressedKeys := []KeyCode{}
	errorTextChannel := GetErrorChannel()
	errorText := ""

	var cancelResetErrorText context.CancelFunc
	resetErrorText := func(ctx context.Context) {
		select {
		case <-time.After(3 * time.Second):
			errorTextChannel <- ""
		case <-ctx.Done():
		}
	}

	for {
		select {
		case activeWindow = <-activeWindowChannel:
		case pressedKeys = <-pressedKeysChannel:
		case errorText = <-errorTextChannel:
			if cancelResetErrorText != nil {
				cancelResetErrorText()
			}
			var ctx context.Context
			ctx, cancelResetErrorText = context.WithCancel(context.Background())
			go resetErrorText(ctx)
		}

		filtered := FilterByApplications(combinations, []string{"windows", activeWindow})
		transformedKeyCombinations := ToKeyCombinationViews(filtered, pressedKeys)
		sortedKeyCombinations := SortByPressedKeys(transformedKeyCombinations)

		ConsoleWriteApplicationState(ApplicationState{
			ActiveWindowTitle:          activeWindow,
			AllKeyCombinations:         combinations,
			PressedKeys:                pressedKeys,
			TransformedKeyCombinations: sortedKeyCombinations,
			Error:                      errorText,
		})

	}
}
