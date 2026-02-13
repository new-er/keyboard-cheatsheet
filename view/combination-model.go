package view

import (
	"fmt"
	"new-er/keyboard-cheatsheet/data"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type combinationModel struct {
	Combination data.Combination
	Keys        nodeModel
}

func (m combinationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	Keys, _ := m.Keys.Update(msg)
	return combinationModel{
		Combination: m.Combination,
		Keys:        Keys,
	}, nil
}

func (m combinationModel) View() string {
	s := fmt.Sprintf("%v [%v] - %v", m.Combination.Description, strings.Join(m.Combination.Applications, ", "), m.Keys.View())
	return s
}

func InitialCombinationModel(combination data.Combination) combinationModel {
	return combinationModel{Combination: combination, Keys: InitialNodeModel(combination.Keys)}
}

func (m combinationModel) Init() tea.Cmd {
	return nil
}

func SortCombinations(combinations []combinationModel, prioritizedApplications []string, searchedString string) []combinationModel {
	slices.SortFunc(combinations, func(a combinationModel, b combinationModel) int {
		matchingApplicationsA := countMatchingApplications(a, prioritizedApplications)
		matchingApplicationsB := countMatchingApplications(b, prioritizedApplications)

		matchingKeysA := a.Keys.GetActiveKeys()
		matchingKeysB := b.Keys.GetActiveKeys()

		compareA := matchingApplicationsA*10000 + matchingKeysA * 10 
		compareB := matchingApplicationsB*10000 + matchingKeysB * 10 

		textA := strings.ToLower(a.View())
		textB := strings.ToLower(b.View())
		matchingTextA := strings.Count(textA, strings.ToLower(searchedString))
		matchingTextB := strings.Count(textB, strings.ToLower(searchedString))
		matchingTextComparison := (matchingTextB - matchingTextA) * 1000000
		if searchedString == "" {
			matchingTextComparison = 0
		}

		return matchingTextComparison + compareB - compareA //return compareB - compareA + textComparison 
	})
	slices.Reverse(combinations)
	return combinations
}

func countMatchingApplications(combination combinationModel, prioritizedApplications []string) int {
	matchingApplications := 0
	for _, application := range combination.Combination.Applications {
		for _, prioritizedApplication := range prioritizedApplications {
			if strings.Contains(strings.ToLower(prioritizedApplication), strings.ToLower(application)) {
				matchingApplications++
			}
		}
	}
	return matchingApplications
}
