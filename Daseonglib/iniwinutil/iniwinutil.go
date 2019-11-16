package iniwinutil

import (
	"syscall"
	"unsafe"
)

var (
	kernel32A                     = syscall.NewLazyDLL("kernel32.dll")
	procGetPrivateProfileString   = kernel32A.NewProc("GetPrivateProfileStringW")
	procWritePrivateProfileString = kernel32A.NewProc("WritePrivateProfileStringW")
)

func GetProfileString(lpszSection string, lpKeyName string, lpDefault string, lpFilePath string) string {
	var ptrValue [2048]uint16
	var ptrlpszSection uintptr
	var ptrlpKeyName uintptr
	var ptrlpDefault uintptr
	var ptrlpFilePath uintptr

	if len(lpszSection) != 0 {
		ptrlpszSection = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpszSection)))
	} else {
		return ""
	}

	if len(lpKeyName) != 0 {
		ptrlpKeyName = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpKeyName)))
	} else {
		return ""
	}

	ptrlpDefault = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpDefault)))

	if len(lpFilePath) != 0 {
		ptrlpFilePath = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpFilePath)))
	} else {
		return ""
	}

	ret, _, _ := procGetPrivateProfileString.Call(ptrlpszSection, ptrlpKeyName, ptrlpDefault,
		uintptr(unsafe.Pointer(&ptrValue[0])), uintptr(len(ptrValue)), ptrlpFilePath)

	if ret <= 0 {
		return ""
	}

	return syscall.UTF16ToString(ptrValue[0:ret])
}

func SetProfileString(lpszSection string, lpKeyName string, lpValue string, lpFilePath string) bool {

	var ptrlpszSection uintptr
	var ptrlpKeyName uintptr
	var ptrlpValue uintptr
	var ptrlpFilePath uintptr

	if len(lpszSection) != 0 {
		ptrlpszSection = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpszSection)))
	} else {
		return false
	}

	if len(lpKeyName) != 0 {
		ptrlpKeyName = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpKeyName)))
	} else {
		return false
	}

	if len(lpValue) != 0 {
		ptrlpValue = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpValue)))
	} else {
		return false
	}

	if len(lpFilePath) != 0 {
		ptrlpFilePath = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpFilePath)))
	} else {
		return false
	}

	ret, _, _ := procWritePrivateProfileString.Call(ptrlpszSection, ptrlpKeyName, ptrlpValue, ptrlpFilePath)
	if ret <= 0 {
		return false
	}

	return true
}
