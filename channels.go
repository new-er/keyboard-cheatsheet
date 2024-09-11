package main

import (
	"keyboard-cheatsheet/main/data"
	"reflect"
	"time"
)

func GetPressedKeysChannel() chan []data.KeyCode {
	currentPressedKeys := []data.KeyCode{}
	pressedKeysChannel := make(chan []data.KeyCode)

	go func() {
		for {
			pressedKeys := GetPressedKeys()
			if !reflect.DeepEqual(pressedKeys, currentPressedKeys) {
				pressedKeysChannel <- pressedKeys
				currentPressedKeys = pressedKeys
			}
			time.Sleep(40 * time.Millisecond)
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
			time.Sleep(40 * time.Millisecond)
		}
	}()
	return activeWindowTitleChannel
}

var (
	errorChannel = make(chan string)
)

func WriteError(err string) {
	errorChannel <- err
}
func GetErrorChannel() chan string {
	return errorChannel
}
