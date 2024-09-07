package main

import "strings"

func FilterByApplications(keyCombinations []KeyCombination, applications []string) []KeyCombination {
	return Filter(keyCombinations, func(keyCombination KeyCombination) bool {
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
