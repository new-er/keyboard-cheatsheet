package repository

import "keyboard-cheatsheet/main/data"

type Repository struct {
	KeyCombinations []KeyCombinationRepository
}

func NewRepository(keyCombinations []data.KeyCombination) *Repository {
	keyCombinationsViewModel := make([]KeyCombinationRepository, len(keyCombinations))
	for i, keyCombination := range keyCombinations {
		keyCombinationsViewModel[i] = NewKeyCombinationRepository(keyCombination)
	}
	return &Repository{}
}
