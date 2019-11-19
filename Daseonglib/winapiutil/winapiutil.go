package winapiutil

import (
	"errors"
	_ "fmt"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

const (
	WH_MOUSE_LL           = 14
	WM_CLOSE              = 16
	WM_QUIT               = 18
	WM_LBUTTONDOWN        = 513
	WM_LBUTTONUP          = 514
	WM_RBUTTONDOWN        = 516
	WM_RBUTTONUP          = 517
	WM_USER               = 1024
	PT_GETWINDOWFROMPOINT = WM_USER + 1
)

type POINT struct {
	X, Y int32
}

type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type MSLLHOOKSTRUCT struct {
	pt          POINT
	mouseData   uint32
	flags       uint32
	time        uint32
	dwExtraInfo uintptr
}

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

type (
	HANDLE    uintptr
	HHOOK     HANDLE
	HINSTANCE HANDLE
	DWORD     uint32
	LRESULT   uintptr
	LPARAM    uintptr
	WPARAM    uintptr
	HWND      HANDLE
	BOOL      int32
	CSIDL     uint32
)

const (
	SW_HIDE = 0
	SW_SHOW = 5
)

const (
	MAX_PATH = 260

	CSIDL_PROGRAM_FILES           = 0x26
	CSIDL_PROGRAM_FILESX86        = 0x2A
	CSIDL_DESKTOPDIRECTORY        = 0x10
	CSIDL_COMMON_DESKTOPDIRECTORY = 0x19
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	shell32  = syscall.NewLazyDLL("shell32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	GetModuleFileNameProc = kernel32.NewProc("GetModuleFileNameW")

	procGetPrivateProfileString   = kernel32.NewProc("GetPrivateProfileStringW")
	procWritePrivateProfileString = kernel32.NewProc("WritePrivateProfileStringW")

	procShellExecuteW = shell32.NewProc("ShellExecuteW")

	procSHGetSpecialFolderPathW = shell32.NewProc("SHGetSpecialFolderPathW")

	procFindWindowW  = user32.NewProc("FindWindowW")
	procSendMessageW = user32.NewProc("SendMessageW")
	procPostMessageW = user32.NewProc("PostMessageW")

	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32First           = kernel32.NewProc("Process32FirstW")
	procProcess32Next            = kernel32.NewProc("Process32NextW")
	procCloseHandle              = kernel32.NewProc("CloseHandle")
	procOpenProcess              = kernel32.NewProc("OpenProcess")
	procTerminateProcess         = kernel32.NewProc("TerminateProcess")

	GetModuleHandleWProc = kernel32.NewProc("GetModuleHandleW")
	ExitProcessProc      = kernel32.NewProc("ExitProcess")

	GetMessageWProc          = user32.NewProc("GetMessageW")
	GetWindowTextLengthWProc = user32.NewProc("GetWindowTextLengthW")
	GetWindowTextWProc       = user32.NewProc("GetWindowTextW")
	GetClassNameWProc        = user32.NewProc("GetClassNameW")

	libuser32, _           = syscall.LoadLibrary("user32.dll")
	WindowFromPointProc, _ = syscall.GetProcAddress(libuser32, "WindowFromPoint")
)

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := GetWindowTextLengthWProc.Call(uintptr(hwnd))
	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1
	buf := make([]uint16, textLen)
	GetWindowTextWProc.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(textLen))
	return syscall.UTF16ToString(buf)
}

func GetClassName(hwnd HWND) string {
	b := make([]uint16, 1024)
	p := uintptr(unsafe.Pointer(&b[0]))
	GetClassNameWProc.Call(uintptr(hwnd), p, 1024)
	return string(syscall.UTF16ToString(b))
}

func WindowFromPoint(Point POINT) HWND {
	ret, _, _ := syscall.Syscall(WindowFromPointProc, 2, uintptr(Point.X), uintptr(Point.Y), 0)
	return HWND(ret)
}

func GetModuleHandle(modulename string) HINSTANCE {
	var mn uintptr
	if modulename == "" {
		mn = 0
	} else {
		mn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(modulename)))
	}
	ret, _, _ := GetModuleHandleWProc.Call(mn)
	return HINSTANCE(ret)
}

func ExitProcess(uExitCode uint32) {
	ExitProcessProc.Call(uintptr(uExitCode))
}

func TerminateProcess(pid int) bool {

	handle, _, _ := procOpenProcess.Call(syscall.PROCESS_TERMINATE, uintptr(0), uintptr(pid))
	if handle < 0 {
		return false
	}
	defer procCloseHandle.Call(handle)

	ret, _, _ := procTerminateProcess.Call(handle, uintptr(0))
	if ret != 1 {
		return false
	}
	return true
}

func GetPID(szFileName string) int {

	var PID int

	snapshot, _, _ := procCreateToolhelp32Snapshot.Call(syscall.TH32CS_SNAPPROCESS, 0)
	if snapshot < 0 {
		return 0
	}
	defer procCloseHandle.Call(snapshot)

	var entry syscall.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	ret, _, _ := procProcess32First.Call(snapshot, uintptr(unsafe.Pointer(&entry)))
	if ret < 0 {
		return 0
	}

	for {

		ExeName := strings.ToLower(syscall.UTF16ToString(entry.ExeFile[:]))
		if ExeName == szFileName {
			PID = int(entry.ProcessID)
			//fmt.Printf("%d %d %s \n", entry.ProcessID, entry.ParentProcessID, ExeName)
			break
		}

		ret, _, _ := procProcess32Next.Call(snapshot, uintptr(unsafe.Pointer(&entry)))
		if ret == 0 {
			break
		}
	}

	return PID
}

func FindWindow(className string, windowName string) (HWND, unsafe.Pointer, HANDLE) {

	ptrclassName := syscall.StringToUTF16Ptr(className)
	ptrwindowName := syscall.StringToUTF16Ptr(windowName)

	ret, _, _ := procFindWindowW.Call(uintptr(unsafe.Pointer(ptrclassName)), uintptr(unsafe.Pointer(ptrwindowName)))
	return HWND(ret), unsafe.Pointer(ret), HANDLE(ret)
}

func FindWindowTitle(sTitle string) HWND {
	ret, _, _ := procFindWindowW.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(sTitle))))
	return HWND(ret)
}

func SendMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {

	ret, _, _ := procSendMessageW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return ret
}

func PostMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) bool {

	ret, _, _ := procPostMessageW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return ret != 0
}

func GetModuleFileName() string {

	var wpath [syscall.MAX_PATH]uint16
	r1, _, _ := GetModuleFileNameProc.Call(0, uintptr(unsafe.Pointer(&wpath[0])), uintptr(len(wpath)))
	if r1 == 0 {
		return ""
	}
	return syscall.UTF16ToString(wpath[:])
}

func GetModulePath() string {

	var wpath [syscall.MAX_PATH]uint16
	r1, _, _ := GetModuleFileNameProc.Call(0, uintptr(unsafe.Pointer(&wpath[0])), uintptr(len(wpath)))
	if r1 == 0 {
		return ""
	}
	return filepath.Dir(syscall.UTF16ToString(wpath[:]))
}

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

func ShellExecute(hwnd HWND, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd int) error {

	var ptrlpOperation uintptr
	var ptrlpFile uintptr
	var ptrlpParameters uintptr
	var ptrlpDirectory uintptr

	if len(lpOperation) != 0 {
		ptrlpOperation = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpOperation)))
	} else {
		ptrlpOperation = uintptr(0)
	}

	if len(lpFile) != 0 {
		ptrlpFile = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpFile)))
	} else {
		ptrlpFile = uintptr(0)
	}

	if len(lpParameters) != 0 {
		ptrlpParameters = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpParameters)))
	} else {
		ptrlpParameters = uintptr(0)
	}

	if len(lpDirectory) != 0 {
		ptrlpDirectory = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpDirectory)))
	} else {
		ptrlpDirectory = uintptr(0)
	}

	ret, _, _ := procShellExecuteW.Call(uintptr(hwnd), ptrlpOperation, ptrlpFile, ptrlpParameters, ptrlpDirectory, uintptr(nShowCmd))

	errMsg := ""
	if ret != 0 && ret <= 32 {
		errMsg = "error"
	} else {
		return nil
	}
	return errors.New(errMsg)
}

func GetDesktopPath() string {

	var buf [MAX_PATH]uint16

	if !SHGetSpecialFolderPath(0, &buf[0], CSIDL_DESKTOPDIRECTORY, false) {
		return ""
	}

	return (syscall.UTF16ToString(buf[0:]))
}

func GetProgramFilesPath() string {

	var buf [MAX_PATH]uint16

	if !SHGetSpecialFolderPath(0, &buf[0], CSIDL_PROGRAM_FILES, false) {
		return ""
	}

	return (syscall.UTF16ToString(buf[0:]))
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}
	return 0
}

func SHGetSpecialFolderPath(hwndOwner HWND, lpszPath *uint16, csidl CSIDL, fCreate bool) bool {

	ret, _, _ := procSHGetSpecialFolderPathW.Call(uintptr(hwndOwner), uintptr(unsafe.Pointer(lpszPath)), uintptr(csidl), uintptr(BoolToBOOL(fCreate)))

	return ret != 0
}
