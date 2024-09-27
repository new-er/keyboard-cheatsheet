package main

import (
	"context"
	"image/color"
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/ui"
	"keyboard-cheatsheet/main/view"
	"keyboard-cheatsheet/main/windows"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	combinationsFile = "combinations.json"
)

func main() {
	a := app.NewWithID("key-combinations")
	w := a.NewWindow("Key Combinations")
	w.Resize(fyne.NewSize(800, 600))

	combinations := data.KeyCombinationsFromFileOrPanic(combinationsFile)
	combinations = data.FilterDisabledKeyCombinations(combinations)
	combinationViews := view.ToKeyCombinationViews(combinations)

	activeWindowChannel := windows.GetActiveWindowTitleChannel()
	activeWindow := ""

	pressedKeysChannel := windows.GetPressedKeysChannel()
	pressedKeys := []data.KeyCode{}

	errorTextChannel := windows.GetErrorChannel()

	sortedKeyCombinations := []view.KeyCombinationView{}
	sortedKeyCombinationsList := widget.NewList(
		func() int {
			return len(sortedKeyCombinations)
		},
		func() fyne.CanvasObject {
			hbox := container.NewHBox()
			hbox.Add(ui.NewText(""))
			return hbox
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			combination := sortedKeyCombinations[id]
			hBox := item.(*fyne.Container)
			hBox.RemoveAll()
			for _, key := range combination.Keys {
				text := ui.NewText(key.Key + " ")
				text.Alignment = fyne.TextAlignCenter
				if key.IsPressed() {
					text.Color = color.RGBA{0, 255, 0, 255}
					text.TextStyle = fyne.TextStyle{Bold: true}
				}
				hBox.Add(key.CanvasText())
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
			case <-errorTextChannel:
				if cancelResetErrorText != nil {
					cancelResetErrorText()
				}
				var ctx context.Context
				ctx, cancelResetErrorText = context.WithCancel(context.Background())
				go resetErrorText(ctx)
			}
			filtered := view.FilterByApplications(combinationViews, []string{"Windows", "PowerToys", activeWindow})
			filtered = view.UpdatePressedKeys(filtered, pressedKeys)
			sortedKeyCombinations = view.SortByPressedKeys(filtered)
			sortedKeyCombinationsList.Refresh()
		}
	}()

	content := container.New(
		layout.NewStackLayout(),
    sortedKeyCombinationsList,
	)
	w.SetContent(content)
	w.ShowAndRun()
}
