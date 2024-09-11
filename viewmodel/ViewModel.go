package viewmodel

import (
	"keyboard-cheatsheet/main/repository"
	"keyboard-cheatsheet/main/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ViewModel struct {
	list *widget.List
}

func (vm *ViewModel) GetList() *widget.List {
	return vm.list
}

func NewViewModel(repository *repository.Repository) *ViewModel {
	keyCombinations := make([]KeyCombinationViewModel, len(repository.KeyCombinations))
	for i, keyCombination := range repository.KeyCombinations {
		keyCombinations[i] = *NewKeyCombinationViewModel(keyCombination)
	}
	println("keyCombinations", keyCombinations)
	list := widget.NewList(
		func() int {
			return len(keyCombinations)
		},
		func() fyne.CanvasObject {
			hbox := container.NewHBox()
			hbox.Add(ui.NewText(""))
			return hbox
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			combination := keyCombinations[id]
			hBox := item.(*fyne.Container)
			hBox.RemoveAll()
			hBox.Add(combination.container)
		},
	)

	return &ViewModel{
		list: list,
	}
}
