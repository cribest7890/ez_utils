//go:build windows

package sys

import (
	"runtime"
	"syscall"
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	procGetGlobalMem = modkernel32.NewProc("GlobalMemoryStatusEx")
	modadvapi32      = syscall.NewLazyDLL("advapi32.dll")
	procOpenToken    = modadvapi32.NewProc("OpenProcessToken")
	procGetTokenInfo = modadvapi32.NewProc("GetTokenInformation")
)

type memoryStatusEx struct {
	cbLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func MemoryUsage() (uint64, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	var stat memoryStatusEx
	stat.cbLength = uint32(unsafeSizeof(stat))
	r1, _, _ := procGetGlobalMem.Call(uintptr(syscall.Pointer(&stat)))
	if r1 == 0 {
		return m.Sys / 1024 / 1024, 0
	}
	return m.Sys / 1024 / 1024, stat.ullAvailPhys / 1024 / 1024
}

func IsRoot() bool {
	hProcess, _ := syscall.GetCurrentProcess()
	var hToken syscall.Token
	r1, _, _ := procOpenToken.Call(uintptr(hProcess), 0x0008, uintptr(syscall.Pointer(&hToken)))
	if r1 == 0 {
		return false
	}
	defer hToken.Close()
	var elevation uint32
	var cbElevation uint32
	r1, _, _ = procGetTokenInfo.Call(uintptr(hToken), 20, uintptr(syscall.Pointer(&elevation)), uintptr(4), uintptr(syscall.Pointer(&cbElevation)))
	if r1 == 0 {
		return false
	}
	return elevation != 0
}

func unsafeSizeof(v interface{}) uintptr {
	switch v.(type) {
	case memoryStatusEx:
		return 64
	default:
		return 0
	}
}

