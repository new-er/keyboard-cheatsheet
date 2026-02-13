package data

import (
	"encoding/json"
	"io"
	"new-er/keyboard-cheatsheet/data/ast"
	"new-er/keyboard-cheatsheet/data/ast/nodes"
	"new-er/keyboard-cheatsheet/helpers"
	"os"
	"slices"
)

type Combination struct {
	Keys         nodes.Node
	Applications []string
	Description  string
	OS           []string
	Disabled     bool
}

type CombinationJson struct {
	Keys         string
	Applications []string
	Description  string
	OS           []string
	Disabled     bool
}

func ReadCombinationsFromFile() []Combination {
	var combinationsFile = "combinations.json"
	var file, err = os.Open(combinationsFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var combinations []CombinationJson
	err = json.Unmarshal(byteValue, &combinations)
	if err != nil {
		panic(err)
	}

	return helpers.Map(combinations, func(combination CombinationJson) Combination {
		lexer := ast.NewLexer(combination.Keys)
		tokens := lexer.AllTokens()
		parser := ast.NewParser(tokens)
		keys := parser.Parse()

		return Combination{
			Keys:         keys,
			Applications: combination.Applications,
			Description:  combination.Description,
			OS:           combination.OS,
			Disabled:     combination.Disabled,
		}
	})
}

func (c Combination) Equals(other helpers.Equatable) bool {
	otherCombination, ok := other.(Combination)
	if !ok {
		return false
	}

	return c.Keys.Equals(otherCombination.Keys) &&
		helpers.StringSliceEquals(c.Applications, otherCombination.Applications) &&
		c.Description == otherCombination.Description &&
		helpers.StringSliceEquals(c.OS, otherCombination.OS) &&
		c.Disabled == otherCombination.Disabled
}

func SortCombinationsByApplications(combinations []Combination, prioritizedApplications []string) []Combination {
	slices.SortFunc(combinations, func(a Combination, b Combination) int {
		aMatchingApplications := countMatchingApplications(a, prioritizedApplications)
		bMatchingApplications := countMatchingApplications(b, prioritizedApplications)
		return bMatchingApplications - aMatchingApplications
	})
	return combinations
}

func countMatchingApplications(combination Combination, prioritizedApplications []string) int {
	matchingApplications := 0
	for _, application := range combination.Applications {
		for _, prioritizedApplication := range prioritizedApplications {
			if application == prioritizedApplication {
				matchingApplications++
			}
		}
	}
	return matchingApplications
}
