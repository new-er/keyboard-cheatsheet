package main

import "strings"

func FilterByApplications(keyCombinations []KeyCombination, applications []string) []KeyCombination {
	return Filter(keyCombinations, func(keyCombination KeyCombination) bool {
		for _, application := range applications {
			if strings.Contains(application, keyCombination.Application) {
				return true
			}
		}
		return false
	})
}
