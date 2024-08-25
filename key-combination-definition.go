package main

func NewKeyCombinationDefinition() []ProgramGroup {
	return []ProgramGroup{
		{
			name: "windows",
			keyCombinations: []KeyCombination{

				NewKeyCombination([]KeyCode{ALT, TAB}, "Switch between open apps"),
			},
		},
		{
			name: "Firefox",
			keyCombinations: []KeyCombination{
				NewKeyCombination([]KeyCode{CTRL, T}, "Open a new tab"),
        NewKeyCombination([]KeyCode{CTRL, SHIFT, T}, "Reopen the last closed tab"),
        NewKeyCombination([]KeyCode{CTRL, F4}, "Close the current tab"),
			},
		},
		{
			name: "PowerShell",
			keyCombinations: []KeyCombination{
				NewKeyCombination([]KeyCode{CTRL, SHIFT, TAB}, "Create a new tab"),
			},
		},
	}
}
