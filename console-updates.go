package main

import (
	"fmt"
	"os"
	"os/exec"
)

func WriteLine(text string) {
	fmt.Println(text)
}

func ClearLines() {
	cmd := exec.Command("cmd", "/c", "cls") // Clear console and reset cursor in Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
