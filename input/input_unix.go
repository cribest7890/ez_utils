//go:build !windows

package input

import (
	"os"
	"syscall"
	"unsafe"
)

var origTerm syscall.Termios

func Setup() error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if isSetup {
		return nil
	}

	fd := int(os.Stdin.Fd())
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&origTerm)), 0, 0, 0); err != 0 {
		return err
	}

	raw := origTerm
	raw.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	raw.Oflag &^= syscall.OPOST
	raw.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	raw.Cflag &^= syscall.CSIZE | syscall.PARENB
	raw.Cflag |= syscall.CS8
	raw.Cc[syscall.VMIN] = 1
	raw.Cc[syscall.VTIME] = 0

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETSW), uintptr(unsafe.Pointer(&raw)), 0, 0, 0); err != 0 {
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

	fd := int(os.Stdin.Fd())
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETSW), uintptr(unsafe.Pointer(&origTerm)), 0, 0, 0); err != 0 {
		return err
	}

	isSetup = false
	return nil
}

