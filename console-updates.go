package main

import (
	"fmt"
)

var (
	writtenLines int
)

func WriteLine(text string) {
	fmt.Println(text)
	writtenLines++
}

func ClearLines() {
	for i := 0; i < writtenLines; i++ {
		fmt.Print("\x1b[1A\x1b[2K")
	}
	writtenLines = 0
}
