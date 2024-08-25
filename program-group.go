package main

import "strings"

type ProgramGroup struct {
	name            string
	keyCombinations []KeyCombination
}

func FilterByProgram(programGroups []ProgramGroup, programNames []string) []ProgramGroup {
	filteredProgramGroups := []ProgramGroup{}
	for _, programGroup := range programGroups {
		for _, programName := range programNames {
			if strings.Contains(programName, programGroup.name) {
				filteredProgramGroups = append(filteredProgramGroups, programGroup)
			}
		}
	}
	return filteredProgramGroups
}
