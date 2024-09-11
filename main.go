package main

import (
	"context"
	"fmt"
	"image/color"
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/repository"
	"keyboard-cheatsheet/main/viewmodel"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	combinationsFile = "combinations.json"
)

func main() {
	combinations2 := data.KeyCombinationsFromFileOrPanic(combinationsFile)
	combinations2 = data.FilterDisabledKeyCombinations(combinations2)

	repository := repository.NewRepository(combinations2)
	viewModel := viewmodel.NewViewModel(repository)

	combinations := KeyCombinationsFromFileOrPanic(combinationsFile)
	combinations = FilterDisabledKeyCombinations(combinations)
	activeWindowChannel := GetActiveWindowTitleChannel()
	activeWindow := ""
	activeWindowBinding := binding.BindString(&activeWindow)

	pressedKeysChannel := GetPressedKeysChannel()
	pressedKeys := []KeyCode{}
	pressedKeysString := ""
	pressedKeysStringBinding := binding.BindString(&pressedKeysString)

	errorTextChannel := GetErrorChannel()
	errorText := ""
	errorTextBinding := binding.BindString(&errorText)

	sortedKeyCombinations := []KeyCombinationView{}
	sortedKeyCombinationsList := widget.NewList(
		func() int {
			return len(sortedKeyCombinations)
		},
		func() fyne.CanvasObject {
			hbox := container.NewHBox()
			hbox.Add(NewText(""))
			return hbox
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			combination := sortedKeyCombinations[id]
			hBox := item.(*fyne.Container)
			hBox.RemoveAll()
			for _, key := range combination.keys {
				text := NewText(key.key)
				text.Alignment = fyne.TextAlignCenter
				if key.isPressed {
					text.Color = color.RGBA{0, 255, 0, 255}
					text.TextStyle = fyne.TextStyle{Bold: true}
				}
				hBox.Add(text)
			}
			hBox.Add(layout.NewSpacer())
			hBox.Add(combination.DescriptionText)
			hBox.Add(combination.ApplicationsContainer)
		},
	)

	var cancelResetErrorText context.CancelFunc
	resetErrorText := func(ctx context.Context) {
		select {
		case <-time.After(3 * time.Second):
			errorTextChannel <- ""
		case <-ctx.Done():
		}
	}

	go func() {

		for {
			select {
			case activeWindow = <-activeWindowChannel:
			case pressedKeys = <-pressedKeysChannel:
			case errorText = <-errorTextChannel:
				if cancelResetErrorText != nil {
					cancelResetErrorText()
				}
				var ctx context.Context
				ctx, cancelResetErrorText = context.WithCancel(context.Background())
				go resetErrorText(ctx)
			}

			activeWindowBinding.Set(activeWindow)
			errorTextBinding.Set(errorText)
			pressedKeysStringBinding.Set(ToPressedKeyString(pressedKeys))

			filtered := FilterByApplications(combinations, []string{"Windows", "PowerToys", activeWindow})
			transformedKeyCombinations := ToKeyCombinationViews(filtered, pressedKeys)
			sortedKeyCombinations = SortByPressedKeys(transformedKeyCombinations)
			sortedKeyCombinationsList.Refresh()
		}
	}()

	a := app.NewWithID("key-combinations")
	w := a.NewWindow("Key Combinations")
	w.Resize(fyne.NewSize(800, 600))

	content := container.New(
		layout.NewStackLayout(),
		viewModel.GetList(),
		//sortedKeyCombinationsList,
	)
	w.SetContent(content)
	w.ShowAndRun()
}

func ToPressedKeyString(pressedKeys []KeyCode) string {
	pressedKeysInterface := make([]string, len(pressedKeys))
	for i, key := range pressedKeys {
		pressedKeysInterface[i] = fmt.Sprint("[", string(key), "]")
	}
	return strings.Join(pressedKeysInterface, " ")
}

func ToKeyCombinationViewsInterface(keyCombinationViews []KeyCombinationView) []interface{} {
	interfaces := make([]interface{}, len(keyCombinationViews))
	for i, keyCombinationView := range keyCombinationViews {
		interfaces[i] = keyCombinationView
	}
	return interfaces
}

func NewText(text string) *canvas.Text {
	canvasText := canvas.NewText(text, color.White)
	canvasText.TextSize = 30
	return canvasText
}
