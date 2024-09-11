package main

import (
	"fmt"
	"keyboard-cheatsheet/main/data"
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow") //GetForegroundWindow
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")      //GetWindowTextW
)

func getForegroundWindow() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
		return
	}
	hwnd = syscall.Handle(r0)
	return
}

func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetActiveWindowTitle() string {
	g, _ := getForegroundWindow()
	b := make([]uint16, 200)
	_, err := getWindowText(g, &b[0], int32(len(b)))
	if err != nil {
		WriteError(fmt.Sprintf("get active window error: %s", err))
		return fmt.Sprintf("error: %s", err)
	}
	return syscall.UTF16ToString(b)
}

func GetPressedKeys() []data.KeyCode {
	var keys []data.KeyCode
	for i := 0; i <= 256; i++ {
		event, _, _ := procGetAsyncKeyState.Call(uintptr(i))
		if event == 0 {
			continue
		}
		if event == 32769 {
			keys = append(keys, data.ToKeycodes(i)...)
		} else if event == 32768 {
			keys = append(keys, data.ToKeycodes(i)...)
		} else if event == 1 {
			keys = append(keys, data.ToKeycodes(i)...)
		} else if event != 0 {
			WriteError(fmt.Sprintf("unknown Val %d for key %s", event, data.ToKeycodes(i)))
		}
	}
	return keys
}
