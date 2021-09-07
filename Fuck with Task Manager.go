package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

var (
	user32             = syscall.MustLoadDLL("user32.dll")
	procGetWindowTextW = user32.MustFindProc("GetWindowTextW")
	procEnumWindows    = user32.MustFindProc("EnumWindows")
	procShowWindow     = user32.MustFindProc("ShowWindow")
	procEnumChildWindows   = user32.MustFindProc("EnumChildWindows")

	kernel32    = syscall.MustLoadDLL("kernel32.dll")
	closeHandle = kernel32.MustFindProc("CloseHandle")

	state bool = false
)

//	SW_HIDE            = 0
//	SW_NORMAL          = 1
//	SW_SHOWNORMAL      = 1
//	SW_SHOWMINIMIZED   = 2
//	SW_MAXIMIZE        = 3
//	SW_SHOWMAXIMIZED   = 3
//	SW_SHOWNOACTIVATE  = 4
//	SW_SHOW            = 5
//	SW_MINIMIZE        = 6
//	SW_SHOWMINNOACTIVE = 7
//	SW_SHOWNA          = 8
//	SW_RESTORE         = 9
//	SW_SHOWDEFAULT     = 10
//	SW_FORCEMINIMIZE   = 11

func FuckWithTaskManager(){
	TaskManager := FindWindow("Task Manager")
	if TaskManager != 0 {
		fmt.Println("Task Manager Handle: " + strconv.FormatInt(int64(TaskManager), 16))

		TaskProcTab := GetChildHandle(TaskManager)
		fmt.Println("Task Manager first child: " + strconv.FormatInt(int64(TaskProcTab), 16))

		if state{
			ShowWindow(TaskProcTab, 1)
			state = false
			fmt.Println("Showing")
		}else{
			ShowWindow(TaskProcTab, 0)
			state = true
			fmt.Println("Hiding")
		}
		CloseHandle(TaskProcTab)
		CloseHandle(TaskManager)
	}
}


func main() {
	Menu:
		fmt.Println("")
	fmt.Println("Fuck with Task Manager [Press Anykey] ")
	CommandScan := bufio.NewScanner(os.Stdin)
	CommandScan.Scan()
	FuckWithTaskManager()
	goto Menu

}

func GetChildHandle(hWnd syscall.Handle) syscall.Handle {
	var hndl syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		hndl = h
		return 0
	})
	EnumChildWindows(hWnd, cb, 0)
	return hndl
}

func ShowWindow(hWnd syscall.Handle, nCmdShow int32) bool {
	ret, _, _ := syscall.Syscall(procShowWindow.Addr(), 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)

	return ret != 0
}

func CloseHandle(hObject syscall.Handle) bool {
	ret, _, _ := syscall.Syscall(closeHandle.Addr(), 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}

func EnumChildWindows(hWndParent syscall.Handle, lpEnumFunc, lParam uintptr) bool {
	ret, _, _ := syscall.Syscall(procEnumChildWindows.Addr(), 3,
		uintptr(hWndParent),
		lpEnumFunc,
		lParam)

	return ret != 0
}

func GetWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func FindWindow(title string) syscall.Handle {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := GetWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			return 1
		}
		if strings.Contains(syscall.UTF16ToString(b), title) {
			hwnd = h
			return 0
		}
		return 1
	})
	_ = EnumWindows(cb, 0)
	if hwnd == 0 {
		return 0
	}
	return hwnd
}

func EnumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

