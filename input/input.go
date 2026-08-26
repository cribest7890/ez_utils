package input

import (
	"strings"
	"sync"
)

type Key string

const (
	KeyArrowUp    Key = "up"
	KeyArrowDown  Key = "down"
	KeyArrowRight Key = "right"
	KeyArrowLeft  Key = "left"
	KeySpace      Key = "space"
	KeyEnter      Key = "enter"
	KeyEscape     Key = "escape"
	KeyBackspace  Key = "backspace"
	KeyTab        Key = "tab"
)

var (
	stateMutex sync.RWMutex
	isPressing = make(map[Key]bool)
	wasPressed = make(map[Key]bool)
	isSetup    bool
)

func IsPressing(k Key) bool {
	stateMutex.RLock()
	defer stateMutex.RUnlock()
	return isPressing[k]
}

func Pressed(k Key) bool {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	if wasPressed[k] {
		wasPressed[k] = false
		return true
	}
	return false
}

func ResetEvents() {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	for k := range wasPressed {
		wasPressed[k] = false
	}
}

func processBytes(b []byte) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	for k := range isPressing {
		isPressing[k] = false
	}

	if len(b) >= 3 && b[0] == 27 && b[1] == 91 {
		switch b[2] {
		case 65:
			triggerKey(KeyArrowUp)
		case 66:
			triggerKey(KeyArrowDown)
		case 67:
			triggerKey(KeyArrowRight)
		case 68:
			triggerKey(KeyArrowLeft)
		}
		return
	}

	if len(b) == 1 {
		switch b[0] {
		case 32:
			triggerKey(KeySpace)
		case 13, 10:
			triggerKey(KeyEnter)
		case 27:
			triggerKey(KeyEscape)
		case 127:
			triggerKey(KeyBackspace)
		case 9:
			triggerKey(KeyTab)
		default:
			if b[0] >= 32 && b[0] <= 126 {
				triggerKey(Key(strings.ToLower(string(b[0]))))
			}
		}
	}
}

func triggerKey(k Key) {
	isPressing[k] = true
	wasPressed[k] = true
}

