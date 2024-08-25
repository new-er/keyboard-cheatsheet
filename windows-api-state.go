package main

import "reflect"

type State struct {
	PressedKeys       []KeyCode
	ActiveWindowTitle string
}

var (
	prevState State
)

func getState() State {
	keys := GetPressedKeys()
	title := GetActiveWindowTitle()
	return State{keys, title}
}
func GetChangedStateOrNull() *State {
	currentState := getState()
	if reflect.DeepEqual(prevState, currentState) && prevState.ActiveWindowTitle == currentState.ActiveWindowTitle {
		return nil
	}
	prevState = currentState
	return &currentState
}
