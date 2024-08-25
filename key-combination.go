package main

import (
	"sort"
)

type KeyCombination struct {
	keys        []KeyCode
	description string
}

func NewKeyCombination(keys []KeyCode, description string) KeyCombination {
	return KeyCombination{
		keys:        sortKeys(keys),
		description: description,
	}
}

type KeyCombinationMatchingPair struct {
	keyCombination KeyCombination
	matching       int
}

func SortByPressedKeys(keyCombinations []KeyCombination, pressedKeys []KeyCode) []KeyCombination {
	pairs := Map(keyCombinations, func(keyCombination KeyCombination) KeyCombinationMatchingPair {
		return KeyCombinationMatchingPair{
			keyCombination: keyCombination,
			matching:       getMatching(keyCombination.keys, pressedKeys),
		}
	})
	if len(pressedKeys) > 0 {
		pairs = Filter(pairs, func(pair KeyCombinationMatchingPair) bool {
			return pair.matching > 0
		})
	}
	pairs = Filter(pairs, func(pair KeyCombinationMatchingPair) bool {
		return pair.matching > 0
	})

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].matching > pairs[j].matching
	})

	return Map(pairs, func(pair KeyCombinationMatchingPair) KeyCombination {
		return pair.keyCombination
	})
}

func getMatching(keys1 []KeyCode, keys2 []KeyCode) int {
	matching := 0
	for i, key1 := range keys1 {
		if i >= len(keys2) {
			break
		}
		if key1 == keys2[i] {
			matching++
		}
	}
	return matching
}

func sortKeys(keys []KeyCode) []KeyCode {
	keysCopy := make([]KeyCode, len(keys))
	copy(keysCopy, keys)
	less := func(i, j int) bool {
		return i < j
	}
	sort.Slice(keysCopy, less)
	return keysCopy
}
