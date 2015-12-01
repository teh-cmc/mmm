// +build !darwin

package mmm

import "syscall"

const (
	mmapFlags = syscall.MAP_PRIVATE | syscall.MAP_ANONYMOUS
)
