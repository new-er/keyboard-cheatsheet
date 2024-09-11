package repository

import (
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/observables"
)

type KeyCodeRepository struct {
	Key       data.KeyCode
	IsPressed observables.Observable[bool]
}

func NewKeyCodeRepository(key data.KeyCode) KeyCodeRepository {
	return KeyCodeRepository{
		Key:       key,
		IsPressed: *observables.NewObservable(false),
	}
}
