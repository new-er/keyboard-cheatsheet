package main

func NewKeyCombinationDefinition() []KeyCombination {
	return []KeyCombination{
    NewKeyCombination([]KeyCode{CTRL, T}, "Open a new tab", "Firefox"),
    NewKeyCombination([]KeyCode{CTRL, SHIFT, T}, "Reopen the last closed tab", "Firefox"),
    NewKeyCombination([]KeyCode{CTRL, F4}, "Close the current tab", "Firefox"),
    NewKeyCombination([]KeyCode{CTRL, SHIFT, TAB}, "Create a new tab", "PowerShell"),
    NewKeyCombination([]KeyCode{ALT, TAB}, "Switch between open apps", "windows"),
	}
}
