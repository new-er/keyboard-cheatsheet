package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	repository := GetRepository()
	repository.keyCombinations.SetValue(NewKeyCombinationDefinition())

	go updatePressedKeys()
	go updateActiveWindowTitle()
	updateConsoleOutput()

	fmt.Println("Starting keyboard cheatsheet...")
}

func updateConsoleOutput() {
	repository := GetRepository()
	keyCombinationsChannel := repository.keyCombinations.Subscribe()
	currentApplicationChannel := repository.currentApplication.Subscribe()
	pressedKeysChannel := repository.pressedKeys.Subscribe()

	for {
		select {
		case keyCombinations := <-keyCombinationsChannel:
			ConsoleWriteApplicationState(ApplicationState{ActiveWindowTitle: repository.currentApplication.GetValue(), AllKeyCombinations: keyCombinations, PressedKeys: repository.pressedKeys.GetValue()})
		case currentApplication := <-currentApplicationChannel:
			ConsoleWriteApplicationState(ApplicationState{ActiveWindowTitle: currentApplication, AllKeyCombinations: repository.keyCombinations.GetValue(), PressedKeys: repository.pressedKeys.GetValue()})
		case pressedKeys := <-pressedKeysChannel:
			ConsoleWriteApplicationState(ApplicationState{ActiveWindowTitle: repository.currentApplication.GetValue(), AllKeyCombinations: repository.keyCombinations.GetValue(), PressedKeys: pressedKeys})
		}
	}
}

func updatePressedKeys() {
	repository := GetRepository()
	for {
		pressedKeys := GetPressedKeys()
		currentPressedKeys := repository.pressedKeys.GetValue()
		if !reflect.DeepEqual(pressedKeys, currentPressedKeys) {
			repository.pressedKeys.SetValue(pressedKeys)
    }
		time.Sleep(100)
	}
}

func updateActiveWindowTitle() {
	repository := GetRepository()
	for {
		activeWindowTitle := GetActiveWindowTitle()
		currentActiveWindowTitle := repository.currentApplication.GetValue()
		if activeWindowTitle != currentActiveWindowTitle {
			repository.currentApplication.SetValue(activeWindowTitle)
		}
		time.Sleep(100)
	}
}
