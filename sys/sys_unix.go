//go:build !windows

package sys

import (
	"os"
	"runtime"
	"syscall"
)

func MemoryUsage() (uint64, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	var sysInfo syscall.Sysinfo_t
	if err := syscall.Sysinfo(&sysInfo); err != nil {
		return m.Sys / 1024 / 1024, 0
	}
	freeMem := sysInfo.Freeram * uint64(sysInfo.Unit) / 1024 / 1024
	return m.Sys / 1024 / 1024, freeMem
}

func IsRoot() bool {
	return os.Getuid() == 0
}
