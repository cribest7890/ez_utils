//go:build !windows

package input

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix" // Pacchetto standard per le costanti di sistema Unix
)

// Cambiamo il tipo in unix.Termios per compatibilità cross-platform
var origTerm unix.Termios

func Setup() error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if isSetup {
		return nil
	}

	fd := int(os.Stdin.Fd())
	
	// Usiamo unix.SYS_IOCTL e unix.TCGETS per supportare Linux
	if _, _, err := syscall.Syscall6(unix.SYS_IOCTL, uintptr(fd), uintptr(unix.TCGETS), uintptr(unsafe.Pointer(&origTerm)), 0, 0, 0); err != 0 {
		return err
	}

	raw := origTerm
	// Modifichiamo i flag usando le costanti standard del pacchetto unix
	raw.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	raw.Oflag &^= unix.OPOST
	raw.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	raw.Cflag &^= unix.CSIZE | unix.PARENB
	raw.Cflag |= unix.CS8
	raw.Cc[unix.VMIN] = 1
	raw.Cc[unix.VTIME] = 0

	// Usiamo unix.TCSETSW per applicare le modifiche dopo che l'output è stato trasmesso
	if _, _, err := syscall.Syscall6(unix.SYS_IOCTL, uintptr(fd), uintptr(unix.TCSETSW), uintptr(unsafe.Pointer(&raw)), 0, 0, 0); err != 0 {
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
	// Ripristiniamo lo stato originale usando unix.TCSETSW
	if _, _, err := syscall.Syscall6(unix.SYS_IOCTL, uintptr(fd), uintptr(unix.TCSETSW), uintptr(unsafe.Pointer(&origTerm)), 0, 0, 0); err != 0 {
		return err
	}

	isSetup = false
	return nil
}
