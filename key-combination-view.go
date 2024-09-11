package main

import (
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/linq"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type KeyCodeView struct {
	key       string
	isPressed bool
}
type KeyCombinationView struct {
	keys         []KeyCodeView
	description  string
	Applications []string

	ApplicationsContainer *fyne.Container
	DescriptionText       *canvas.Text
}

func ToKeyCombinationViews(keyCombinations []data.KeyCombination, pressedKeys []data.KeyCode) []KeyCombinationView {
	return linq.Map(keyCombinations, func(keyCombination data.KeyCombination) KeyCombinationView {
		return ToKeyCombinationView(keyCombination, pressedKeys)
	})
}

func ToKeyCombinationView(keyCombination data.KeyCombination, pressedKeys []data.KeyCode) KeyCombinationView {
	keys := linq.Map(keyCombination.Keys, func(key data.KeyCode) KeyCodeView {
		return KeyCodeView{
			key:       string(key),
			isPressed: Contains(pressedKeys, key),
		}
	})
	keys = SortByPressedKeysCode(keys)

	return KeyCombinationView{
		keys:                  keys,
		description:           keyCombination.Description,
		Applications:          keyCombination.Applications,
		DescriptionText:       NewText(keyCombination.Description),
		ApplicationsContainer: NewApplications(keyCombination.Applications),
	}
}

func SortByPressedKeysCode(views []KeyCodeView) []KeyCodeView {
	sort.Slice(views, func(i, j int) bool {
		return views[i].isPressed && !views[j].isPressed
	})

	return views

}

func SortByPressedKeys(views []KeyCombinationView) []KeyCombinationView {
	sort.Slice(views, func(i, j int) bool {
		countI := linq.Count(views[i].keys, func(key KeyCodeView) bool {
			return key.isPressed
		})
		countJ := linq.Count(views[j].keys, func(key KeyCodeView) bool {
			return key.isPressed
		})
		return countI > countJ
	})

	return views
}

func ToText(keyCodes []KeyCodeView) string {
	return strings.Join(linq.Map(keyCodes, func(key KeyCodeView) string {
		return string(key.key)
	}), " + ")
}
func Contains(slice []data.KeyCode, item data.KeyCode) bool {
	for _, key := range slice {
		if item.Matches(key) {
			return true
		}
	}

	return false
}

func NewApplications(applications []string) *fyne.Container {
	hBox := container.NewHBox()
	for _, application := range applications {
		hBox.Add(GetImageOrText(application))
	}
	return hBox
}
