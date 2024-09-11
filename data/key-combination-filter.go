package data

import (
	"keyboard-cheatsheet/main/linq"
	"strings"
)

func FilterDisabledKeyCombinations(k []KeyCombination) []KeyCombination {
	return linq.Filter(k, func(keyCombination KeyCombination) bool {
		return !keyCombination.Disabled
	})
}

func FilterByApplications(keyCombinations []KeyCombination, applications []string) []KeyCombination {
	return linq.Filter(keyCombinations, func(keyCombination KeyCombination) bool {
		for _, application := range applications {
			for _, keyCombinationApplication := range keyCombination.Applications {
				if strings.Contains(application, keyCombinationApplication) {
					return true
				}
			}
		}
		return false
	})
}
