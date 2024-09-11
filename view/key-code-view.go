package view

import (
	"keyboard-cheatsheet/main/data"
	"keyboard-cheatsheet/main/linq"
	"strings"
)

type KeyCodeView struct {
	Key       string
	IsPressed bool
}

func NewKeyCodeView(key string, isPressed bool) KeyCodeView {
	return KeyCodeView{
		Key:       key,
		IsPressed: isPressed,
	}
}
func ToText(keyCodes []KeyCodeView) string {
	return strings.Join(linq.Map(keyCodes, func(key KeyCodeView) string {
		return string(key.Key)
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
