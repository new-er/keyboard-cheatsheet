package view

import (
	"keyboard-cheatsheet/main/linq"
	"sort"
)

func SortByPressedKeysCode(views []KeyCodeView) []KeyCodeView {
	sort.Slice(views, func(i, j int) bool {
		return views[i].IsPressed() && !views[j].IsPressed()
	})

	return views

}

func SortByPressedKeys(views []KeyCombinationView) []KeyCombinationView {
	sort.Slice(views, func(i, j int) bool {
		countI := linq.Count(views[i].Keys, func(key KeyCodeView) bool {
			return key.IsPressed()
		})
		countJ := linq.Count(views[j].Keys, func(key KeyCodeView) bool {
			return key.IsPressed()
		})
		return countI > countJ
	})

	return views
}
