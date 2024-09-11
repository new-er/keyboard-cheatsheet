package view

import (
	"keyboard-cheatsheet/main/linq"
	"strings"
)

func FilterByApplications(keyCombinations []KeyCombinationView, applications []string) []KeyCombinationView {
	return linq.Filter(keyCombinations, func(keyCombination KeyCombinationView) bool {
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
