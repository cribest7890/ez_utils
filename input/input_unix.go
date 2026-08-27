//go:build !windows

package input

import (
	"os"

	"golang.org/x/sys/unix" // Gestisce l'interfaccia con il kernel Linux in modo nativo e sicuro
)

var origTerm unix.Termios

func Setup() error {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if isSetup {
		return nil
	}

	fd := int(os.Stdin.Fd())
	
	// Sostituiamo la Syscall6 con il wrapper nativo di Go per leggere lo stato (TCGETS)
	t, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return err
	}
	origTerm = *t

	raw := origTerm
	
	// Applichiamo i tuoi flag per la modalità Raw
	raw.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	raw.Oflag &^= unix.OPOST
	raw.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	raw.Cflag &^= unix.CSIZE | unix.PARENB
	raw.Cflag |= unix.CS8
	raw.Cc[unix.VMIN] = 1
	raw.Cc[unix.VTIME] = 0

	// Sostituiamo la Syscall6 con il wrapper nativo per applicare i cambiamenti (TCSETSW)
	if err := unix.IoctlSetTermios(fd, unix.TCSETSW, &raw); err != nil {
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
	
	// Ripristiniamo lo stato originale in modo sicuro
	if err := unix.IoctlSetTermios(fd, unix.TCSETSW, &origTerm); err != nil {
		return err
	}

	isSetup = false
	return nil
}
