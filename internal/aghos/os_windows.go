//go:build windows
// +build windows

package aghos

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func canBindPrivilegedPorts() (can bool, err error) {
	return HaveAdminRights()
}

func setRlimit(val uint64) (err error) {
	return Unsupported("setrlimit")
}

func haveAdminRights() (bool, error) {
	var token windows.Token
	h := windows.CurrentProcess()
	err := windows.OpenProcessToken(h, windows.TOKEN_QUERY, &token)
	if err != nil {
		return false, err
	}

	info := make([]byte, 4)
	var returnedLen uint32
	err = windows.GetTokenInformation(token, windows.TokenElevation, &info[0], uint32(len(info)), &returnedLen)
	token.Close()
	if err != nil {
		return false, err
	}
	if info[0] == 0 {
		return false, nil
	}
	return true, nil
}

func sendProcessSignal(pid int, sig syscall.Signal) error {
	return Unsupported("kill")
}

func isOpenWrt() (ok bool) {
	return false
}
