package main

import (
	"fmt"
)

type ApplicationState struct {
  ActiveWindowTitle string
  AllKeyCombinations   []KeyCombination
  PressedKeys       []KeyCode
}

func ConsoleWriteApplicationState(state ApplicationState) {
  ConsoleClearWrittenLines()
  ConsoleWriteLine(state.ActiveWindowTitle)
  ConsoleWriteLine(fmt.Sprint(state.PressedKeys))

  filtered := FilterByApplications(state.AllKeyCombinations, []string{"windows", state.ActiveWindowTitle})
  isPressedKeyCombinations := ToIsPressedKeyCombinations(filtered, state.PressedKeys)
  sorted := SortByPressedKeys(isPressedKeyCombinations)
  for _, keyCombination := range sorted {
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

    keyCombinationKeys = append(keyCombinationKeys, fmt.Sprint(color,key.key,colorReset))
  }

  text := fmt.Sprintf("%s: %s (%s)", keyCombinationKeys, keyCombinationDescription, keyCombinationApplication)
  ConsoleWriteLine(text)
}


func Contains(slice []KeyCode, item KeyCode) bool {
	for _, key := range slice {
		if key == item {
			return true
    }
	}
	return false
}
