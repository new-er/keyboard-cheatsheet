package main

import (
	"sort"
	"strings"
)

type KeyCombination struct {
	keys        []KeyCode
	description string
	application string
}

func NewKeyCombination(keys []KeyCode, description string, application string) KeyCombination {
	return KeyCombination{
		keys:        keys,
		description: description,
		application: application,
	}
}

func FilterByApplications(keyCombinations []KeyCombination, applications []string) []KeyCombination {
	return Filter(keyCombinations, func(keyCombination KeyCombination) bool {
		for _, application := range applications {
			if strings.Contains(application, keyCombination.application) {
				return true
			}
		}
		return false
	})
}

type IsPressedKeyCode struct {
	key       KeyCode
	isPressed bool
}

type IsPressedKeyCombination struct {
	keys        []IsPressedKeyCode
	description string
	application string
}

func ToIsPressedKeyCombinations(keyCombinations []KeyCombination, pressedKeys []KeyCode) []IsPressedKeyCombination {
	return Map(keyCombinations, func(keyCombination KeyCombination) IsPressedKeyCombination {
		return ToIsPressedKeyCombination(keyCombination, pressedKeys)
	})
}

func ToIsPressedKeyCombination(keyCombination KeyCombination, pressedKeys []KeyCode) IsPressedKeyCombination {
	keys := Map(keyCombination.keys, func(key KeyCode) IsPressedKeyCode {
		return IsPressedKeyCode{
			key:       key,
			isPressed: Contains(pressedKeys, key),
		}
	})

	return IsPressedKeyCombination{
		keys:        keys,
		description: keyCombination.description,
		application: keyCombination.application,
	}
}

func SortByPressedKeys(isPressedKeyCombinations []IsPressedKeyCombination) []IsPressedKeyCombination {
	sort.Slice(isPressedKeyCombinations, func(i, j int) bool {
		return len(isPressedKeyCombinations[i].keys) > len(isPressedKeyCombinations[j].keys)
	})

	return isPressedKeyCombinations
}
