package main

import (
	"reflect"
)

func GetPressedKeysChannel() chan []KeyCode {
	currentPressedKeys := []KeyCode{}
	pressedKeysChannel := make(chan []KeyCode)

	go func() {
		for {
			pressedKeys := GetPressedKeys()
			if !reflect.DeepEqual(pressedKeys, currentPressedKeys) {
				pressedKeysChannel <- pressedKeys
				currentPressedKeys = pressedKeys
			}
		}
	}()
	return pressedKeysChannel
}

func GetActiveWindowTitleChannel() chan string {
	currentActiveWindowTitle := ""
	activeWindowTitleChannel := make(chan string)

	go func() {
		for {
			activeWindowTitle := GetActiveWindowTitle()
			if activeWindowTitle != currentActiveWindowTitle {
				activeWindowTitleChannel <- activeWindowTitle
				currentActiveWindowTitle = activeWindowTitle
			}
		}
	}()
	return activeWindowTitleChannel
}

var (
	errorChannel = make(chan string)
)

func PublishError(err string) {
	errorChannel <- err
}
func GetErrorChannel() chan string {
	return errorChannel
}
