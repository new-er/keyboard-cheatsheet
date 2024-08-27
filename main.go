package main

import (
	"context"
	"fmt"
	"image/color"
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
	combinations := KeyCombinationsFromFileOrPanic(combinationsFile)
	activeWindowChannel := GetActiveWindowTitleChannel()
	activeWindow := ""
	activeWindowBinding := binding.BindString(&activeWindow)
	//activeWindowLabel := widget.NewLabelWithData(binding.StringToStringWithFormat(activeWindowBinding, "activeWindow: %s"))

	pressedKeysChannel := GetPressedKeysChannel()
	pressedKeys := []KeyCode{}
	pressedKeysString := ""
	pressedKeysStringBinding := binding.BindString(&pressedKeysString)
	//pressedKeysStringLabel := widget.NewLabelWithData(binding.StringToStringWithFormat(pressedKeysStringBinding, "pressedKeys: %s"))

	errorTextChannel := GetErrorChannel()
	errorText := ""
	errorTextBinding := binding.BindString(&errorText)
	//errorTextLabel := widget.NewLabelWithData(binding.StringToStringWithFormat(errorTextBinding, "errorText: %s"))

	sortedKeyCombinations := []KeyCombinationView{}
	// Create a new list widget.
	sortedKeyCombinationsList := widget.NewList(
		// Length function returns the number of items.
		func() int {
			return len(sortedKeyCombinations)
		},
		// CreateItem function returns the widget for each list item.
		func() fyne.CanvasObject {
			// Here, you create the template for list items, which can be updated in UpdateItem.
			hbox := container.NewHBox()
			hbox.Add(widget.NewLabel(""))
			return hbox
		},
		// UpdateItem function updates the content of each list item based on its index.
		func(id widget.ListItemID, item fyne.CanvasObject) {
			combination := sortedKeyCombinations[id]
			hBox := item.(*fyne.Container)
			hBox.RemoveAll()
			for _, key := range combination.keys {
				text := canvas.NewText(key.key, color.White)
				text.Alignment = fyne.TextAlignCenter
				if key.isPressed {
					text.Color = color.RGBA{0, 255, 0, 255}
					text.TextStyle = fyne.TextStyle{Bold: true}
				}
				hBox.Add(text)
			}
			hBox.Add(layout.NewSpacer())
			hBox.Add(widget.NewLabel(combination.description))
			hBox.Add(widget.NewLabel(combination.Application))
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
		sortedKeyCombinationsList,
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
