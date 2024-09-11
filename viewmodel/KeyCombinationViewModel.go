package viewmodel

import (
	"keyboard-cheatsheet/main/repository"
	"keyboard-cheatsheet/main/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type KeyCombinationViewModel struct {
	container *fyne.Container
}

func NewKeyCombinationViewModel(combination repository.KeyCombinationRepository) *KeyCombinationViewModel {
	container := container.NewHBox()
	for _, key := range combination.Keys {
		key := *NewKeyCodeViewModel(&key)
		container.Add(key.text)
	}

	container.Add(layout.NewSpacer())
	container.Add(ui.NewText(combination.Description))
	container.Add(NewApplications(combination.Applications))
	return &KeyCombinationViewModel{
		container: container,
	}
}
