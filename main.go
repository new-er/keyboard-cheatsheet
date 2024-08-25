package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Starting keyboard cheatsheet...")
	keyCombinations := NewKeyCombinationDefinition()
	fmt.Println("Key combinations:", keyCombinations)

	state := GetChangedStateOrNull()

	for {
		state = GetChangedStateOrNull()
		if state == nil {
			continue
		}

		fmt.Println("ActiveWindowTitle:", state.ActiveWindowTitle)
		fmt.Println("PressedKeys:", state.PressedKeys)

		filtered := FilterByProgram(keyCombinations, []string{"windows", state.ActiveWindowTitle})
		fmt.Println("Filtered by program:", filtered)

		flattened := Flatten(filtered, func(group ProgramGroup) []KeyCombination {
			return group.keyCombinations
		})

		sorted := SortByPressedKeys(flattened, state.PressedKeys)
		fmt.Println("sorted by pressed keys:", sorted)
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
