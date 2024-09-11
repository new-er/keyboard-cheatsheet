package repository

import "keyboard-cheatsheet/main/data"

type KeyCombinationRepository struct {
	Keys         []KeyCodeRepository
	Description  string
	Applications []string
}

func NewKeyCombinationRepository(keyCombination data.KeyCombination) KeyCombinationRepository {
	keys := make([]KeyCodeRepository, len(keyCombination.Keys))
	for i, key := range keyCombination.Keys {
		keys[i] = NewKeyCodeRepository(key)
	}
	return KeyCombinationRepository{
		Keys:         keys,
		Description:  keyCombination.Description,
		Applications: keyCombination.Applications,
	}
}
