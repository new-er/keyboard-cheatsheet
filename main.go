package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Starting keyboard cheatsheet...")

  state := GetChangedStateOrNull()

	for {
    state = GetChangedStateOrNull()
    if state != nil {
      fmt.Println("ActiveWindowTitle:", state.ActiveWindowTitle)
      fmt.Println("PressedKeys:", state.PressedKeys)
    }
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
