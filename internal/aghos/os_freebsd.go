//go:build freebsd
// +build freebsd

package aghos

import (
	"os"
	"syscall"
)

func canBindPrivilegedPorts() (can bool, err error) {
	return HaveAdminRights()
}

func setRlimit(val uint64) (err error) {
	var rlim syscall.Rlimit
	rlim.Max = int64(val)
	rlim.Cur = int64(val)

	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
}

func haveAdminRights() (bool, error) {
	return os.Getuid() == 0, nil
}

func sendProcessSignal(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}

func isOpenWrt() (ok bool) {
	return false
}
