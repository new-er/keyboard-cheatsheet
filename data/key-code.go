package data

import (
	"fmt"
	"keyboard-cheatsheet/main/log"
	"strings"

	"github.com/TheTitanrain/w32"
)

type KeyCode string

const (
	CTRL              KeyCode = "CTRL"
	BACK              KeyCode = "BACK"
	TAB               KeyCode = "TAB"
	ENTER             KeyCode = "⏎"
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
	SEMICOLON         KeyCode = ";"
	SLASH             KeyCode = "/"
	BACKTICK          KeyCode = "`"
	LEFTBRACKET       KeyCode = "{"
	BACKSLASH         KeyCode = "\\"
	RIGHTBRACKET      KeyCode = "}"
	SINGLEQUOTE       KeyCode = "'"
	PERIOD            KeyCode = "."
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
	LBUTTON           KeyCode = "L🖱"
	RBUTTON           KeyCode = "R🖱"
	MBUTTON           KeyCode = "M🖱"
	ARROWUP           KeyCode = "↑"
	ARROWDOWN         KeyCode = "↓"
	ARROWLEFT         KeyCode = "←"
	ARROWRIGHT        KeyCode = "→"
	COMMA             KeyCode = "COMMA"
	EQUALS            KeyCode = "EQUALS"
	HOME              KeyCode = "HOME"
	WINDOWS           KeyCode = "WIN"
	UNKNOWN           KeyCode = "UNKNOWN"
)

func ToKeycode(key int) KeyCode {
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
	case 8:
		return BACK
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
	case 32:
		return SPACE
	case 33:
		return PAGEUP
	case 34:
		return PAGEDOWN
	case 35:
		return END
	case 36:
		return HOME
	case 38:
		return ARROWUP
	case 37:
		return ARROWLEFT
	case 39:
		return ARROWRIGHT
	case 40:
		return ARROWDOWN
	case 46:
		return DELETE
	case 91:
		return RIGHTWINDOWS
	case 92:
		return RIGHTWINDOWS
	case 112:
		return F1
	case 113:
		return F2
	case 114:
		return F3
	case 115:
		return F4
	case 116:
		return F5
	case 117:
		return F6
	case 118:
		return F7
	case 119:
		return F8
	case 120:
		return F9
	case 121:
		return F10
	case 122:
		return F11
	case 123:
		return F12
	case 160:
		return LSHIFT
	case 187:
		return EQUALS
	case 188:
		return COMMA
	case 189:
		return SUBTRACT
	}
	log.LogError(fmt.Sprintf("unknown key: %d", key))
	return KeyCode(fmt.Sprintf("?(%d)", key))
}

func ToKeycodes(key int) []KeyCode {
	keycodes := []KeyCode{}
	keycode := ToKeycode(key)

	keycodes = append(keycodes, keycode)

	if keycode == LEFTWINDOWS || keycode == RIGHTWINDOWS {
		keycodes = append(keycodes, WINDOWS)
	}
	return keycodes
}

func (k KeyCode) MatchesKeyCode(pressedKey KeyCode) bool {
	kString := string(k)
	pressedKeyString := string(pressedKey)
	return MatchesKeyString(kString, pressedKeyString)
}

func MatchesKeyString(kString string, pressedKeyString string) bool {
	if strings.HasPrefix(kString, "<") && strings.HasSuffix(kString, ">") {
		withoutBrackets := kString[1 : len(kString)-1]
		split := strings.Split(withoutBrackets, "|")
		for _, s := range split {
			if s == pressedKeyString {
				return true
			}
		}
	}
	if strings.HasPrefix(kString, "(") && strings.HasSuffix(kString, ")") {
		withoutBrackets := kString[1 : len(kString)-1]
		return withoutBrackets == pressedKeyString
	}
	return kString == pressedKeyString
}
