package view

import (
	"new-er/keyboard-cheatsheet/data"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	searchModeEnabled       bool
	searchString            string
	focusedApplicationTitle string
	pressedKeys             string
	combinations            []combinationModel
}

func InitialModel() model {
	return model{
		focusedApplicationTitle: "myapp",
		pressedKeys:             "mykeys",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func updateCombinations(combinations []combinationModel, msg tea.Msg) []combinationModel {
	combinationModels := []combinationModel{}
	for _, combination := range combinations {
		update, _ := combination.Update(msg)
		combinationModels = append(combinationModels, update.(combinationModel))
	}
	return combinationModels
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+f":
			m.searchModeEnabled = !m.searchModeEnabled
			if m.searchModeEnabled {
				m.searchString = ""
			}
			m.combinations = updateCombinations(m.combinations, searchedStringChangedMsg{searchString: m.searchString})
			m.combinations = SortCombinations(m.combinations, []string{m.focusedApplicationTitle}, m.searchString)
			return m, nil
		case "backspace":
			if m.searchModeEnabled {
				if len(m.searchString) > 0 {
					m.searchString = m.searchString[:len(m.searchString)-1]
					m.combinations = updateCombinations(m.combinations, searchedStringChangedMsg{searchString: m.searchString})
					m.combinations = SortCombinations(m.combinations, []string{m.focusedApplicationTitle}, m.searchString)
				}
			}
			return m, nil
		case "esc":
			if m.searchModeEnabled {
				m.searchModeEnabled = false
				m.searchString = ""
				m.combinations = updateCombinations(m.combinations, searchedStringChangedMsg{searchString: m.searchString})
				m.combinations = SortCombinations(m.combinations, []string{m.focusedApplicationTitle}, m.searchString)
			}
			return m, nil
		default:
			if m.searchModeEnabled {
				m.searchString += msg.String()
				m.combinations = updateCombinations(m.combinations, searchedStringChangedMsg{searchString: m.searchString})
				m.combinations = SortCombinations(m.combinations, []string{m.focusedApplicationTitle}, m.searchString)
			}
			return m, nil
		}
	case FocusedApplicationTitleChangedMsg:
		{
			m.focusedApplicationTitle = msg.Title
			m.combinations = updateCombinations(m.combinations, msg)
			m.combinations = SortCombinations(m.combinations, []string{m.focusedApplicationTitle}, m.searchString)
			return m, nil
		}
	case PressedKeysChangedMsg:
		{
			stringBuilder := strings.Builder{}
			for _, key := range msg.Keys {
				keyName := "Unknown"
				if len(key.Names) > 0 {
					keyName = key.Names[0]
				}
				stringBuilder.WriteString(keyName)
				stringBuilder.WriteString(" ")
			}
			m.pressedKeys = stringBuilder.String()
			m.combinations = updateCombinations(m.combinations, msg)
			m.combinations = SortCombinations(m.combinations, []string{m.focusedApplicationTitle}, m.searchString)
			return m, nil
		}
	case CombinationsChangedMsg:
		{
			combinationModels := []combinationModel{}
			for _, combination := range msg.Combinations {
				combinationModels = append(combinationModels, InitialCombinationModel(combination))
			}
			updateCombinations(combinationModels, msg)
			m.combinations = combinationModels
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""
	for _, combination := range m.combinations {
		s += combination.View()
		s += "\n"
	}
	s += "\n\n"
	s += "Keyboard Cheatsheet - " + m.focusedApplicationTitle + " | " + m.pressedKeys + ""
	if m.searchModeEnabled {
		s += "Search: " + m.searchString
		s += "\nPress ctrl + f to exit search mode"
	} else {
		s += "\nPress ctrl + f to search"
	}
	s += " | Press ctrl + c to quit"
	return s
}

type FocusedApplicationTitleChangedMsg struct {
	Title string
}

type PressedKeysChangedMsg struct {
	Keys []data.Key
}

type CombinationsChangedMsg struct {
	Combinations []data.Combination
}

type searchedStringChangedMsg struct {
	searchString string
}
