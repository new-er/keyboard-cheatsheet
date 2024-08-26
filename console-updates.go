package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func WriteLine(text string) {
	fmt.Println(text)
}

func ClearLines() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MoveCursorTo(x, y int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setConsoleCursorPosition := kernel32.NewProc("SetConsoleCursorPosition")

	handle := syscall.Handle(uintptr(syscall.Stdout))
	pos := uintptr((y << 16) | x)

	setConsoleCursorPosition.Call(uintptr(handle), pos)
}
