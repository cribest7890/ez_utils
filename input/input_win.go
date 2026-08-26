//go:build windows

package input

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
	origMode           uint32
)

const (
	enableProcessedInput uint32 = 0x0001
	enableLineInput      uint32 = 0x0002
	enableEchoInput      uint32 = 0x0004
	enableWindowInput    uint32 = 0x0008
	enableMouseInput     uint32 = 0x0010
	enableInsertMode     uint32 = 0x0020
	enableQuickEditMode  uint32 = 0x0040
	enableExtendedFlags  uint32 = 0x0080
)

func Setup() error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if isSetup {
		return nil
	}

	handle := syscall.Handle(os.Stdin.Fd())
	r1, _, err := procGetConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&origMode)))
	if r1 == 0 {
		return err
	}

	rawMode := origMode &^ (enableLineInput | enableEchoInput | enableProcessedInput)
	r1, _, err = procSetConsoleMode.Call(uintptr(handle), uintptr(rawMode))
	if r1 == 0 {
		return err
	}

	isSetup = true
	go func() {
		buf := make([]byte, 16)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil || n == 0 {
				break
			}
			
			if n >= 1 && buf[0] == 0 {
				if n >= 2 {
					switch buf[1] {
					case 72:
						processBytes([]byte{27, 91, 65})
						continue
					case 80:
						processBytes([]byte{27, 91, 66})
						continue
					case 77:
						processBytes([]byte{27, 91, 67})
						continue
					case 75:
						processBytes([]byte{27, 91, 68})
						continue
					}
				}
			}
			
			if n >= 1 && buf[0] == 224 {
				if n >= 2 {
					switch buf[1] {
					case 72:
						processBytes([]byte{27, 91, 65})
						continue
					case 80:
						processBytes([]byte{27, 91, 66})
						continue
					case 77:
						processBytes([]byte{27, 91, 67})
						continue
					case 75:
						processBytes([]byte{27, 91, 68})
						continue
					}
				}
			}
			
			processBytes(buf[:n])
		}
	}()

	return nil
}

func Restore() error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if !isSetup {
		return nil
	}

	handle := syscall.Handle(os.Stdin.Fd())
	r1, _, err := procSetConsoleMode.Call(uintptr(handle), uintptr(origMode))
	if r1 == 0 {
		return err
	}

	isSetup = false
	return nil
}

