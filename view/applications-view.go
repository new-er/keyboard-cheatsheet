package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewApplicationsView(applications []string) *fyne.Container {
	hBox := container.NewHBox()
	for _, application := range applications {
		hBox.Add(CachedImageOrTextView(application))
	}
	return hBox
}
