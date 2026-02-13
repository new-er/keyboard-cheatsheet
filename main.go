package main

import (
	"fmt"
	"new-er/keyboard-cheatsheet/data"
	"new-er/keyboard-cheatsheet/helpers"
	osapi "new-er/keyboard-cheatsheet/os"
	"new-er/keyboard-cheatsheet/view"
	"os"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	pressedKeys := []data.Key{}
	keyMap := data.ReadKeyMapFromFile()
	combinations := data.ReadCombinationsFromFile()
	combinations = helpers.Filter(combinations, func(combination data.Combination) bool {
		if len(combination.OS) > 0 && combination.OS[0] == "windows" {
			return true
		}
		return combination.Disabled
	})
	p := tea.NewProgram(view.InitialModel())

	go func() {
		activeWindowTitleChannel := osapi.GetActiveWindowTitleChannel()
		keyEventChannel := osapi.GetKeyEvents()

		p.Send(view.CombinationsChangedMsg{Combinations: combinations})

		for {
			select {
			case activeWindowTitle := <-activeWindowTitleChannel:
				p.Send(view.FocusedApplicationTitleChangedMsg{Title: activeWindowTitle})
			case keyEvent := <-keyEventChannel:
				if keyEvent.State == osapi.PRESSED {
					key, success := keyMap[keyEvent.KeyCode]
					if !success {
						key = data.NewUnknownKey(keyEvent.KeyCode)
					}
					pressedKeys = append(pressedKeys, key)
				} else {
					for i, k := range pressedKeys {
						if k.KeyCode == keyEvent.KeyCode {
							pressedKeys = slices.Delete(pressedKeys, i, i+1)
						}
					}
				}
				p.Send(view.PressedKeysChangedMsg{Keys: pressedKeys})
			}
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
