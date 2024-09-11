package main

import (
	"os"
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

func ToKeyCombinationViews(keyCombinations []KeyCombination, pressedKeys []KeyCode) []KeyCombinationView {
	return Map(keyCombinations, func(keyCombination KeyCombination) KeyCombinationView {
		return ToKeyCombinationView(keyCombination, pressedKeys)
	})
}

func ToKeyCombinationView(keyCombination KeyCombination, pressedKeys []KeyCode) KeyCombinationView {
	keys := Map(keyCombination.Keys, func(key KeyCode) KeyCodeView {
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
		countI := Count(views[i].keys, func(key KeyCodeView) bool {
			return key.isPressed
		})
		countJ := Count(views[j].keys, func(key KeyCodeView) bool {
			return key.isPressed
		})
		return countI > countJ
	})

	return views
}

func ToText(keyCodes []KeyCodeView) string {
	return strings.Join(Map(keyCodes, func(key KeyCodeView) string {
		return string(key.key)
	}), " + ")
}
func Contains(slice []KeyCode, item KeyCode) bool {
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
		hBox.Add(NewApplication(application))
	}
	return hBox
}

func NewApplication(application string) fyne.CanvasObject {
	imagePath := "./icons/" + application + ".png"
	if FileExists(imagePath) {
		image := canvas.NewImageFromFile(imagePath)
		image.FillMode = canvas.ImageFillStretch
		image.SetMinSize(fyne.NewSize(30, 30))
		return image
	}
	return NewText(application)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
