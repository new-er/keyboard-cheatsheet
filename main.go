package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Starting keyboard cheatsheet...")
	keyCombinations := NewKeyCombinationDefinition()
	state := GetChangedStateOrNull()

	for {
		state = GetChangedStateOrNull()
		if state == nil {
			continue
		}
		ConsoleClearWrittenLines()

		ConsoleWriteLine("ActiveWindowTitle: " + state.ActiveWindowTitle)
		ConsoleWriteLine("PressedKeys: " + fmt.Sprint(state.PressedKeys))

		filtered := FilterByApplications(keyCombinations, []string{"windows", state.ActiveWindowTitle})
		//ConsoleWriteLine("Filtered by program: " + fmt.Sprint(filtered))
		sorted := SortByPressedKeys(filtered, state.PressedKeys)
		ConsoleWriteLine(fmt.Sprint(sorted))
	}
}

var (
	prevPressedKeys []KeyCode
)

func getPressedKeysOnChangeOrNull() []KeyCode {
	currentPressedKeys := GetPressedKeys()
	if !reflect.DeepEqual(prevPressedKeys, currentPressedKeys) {
		prevPressedKeys = currentPressedKeys
		return currentPressedKeys
	}
	return nil
}

var (
	prevTitle string
)

func getActiveWindowTitleOnChangeOrEmtpy() string {
	currentTitle := GetActiveWindowTitle()
	if currentTitle == "error" {
		return ""
	}
	if currentTitle == prevTitle {
		return ""
	}
	prevTitle = currentTitle
	return currentTitle
}
