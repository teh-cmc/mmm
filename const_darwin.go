package mmm

import "golang.org/x/sys/unix"

const (
	mmapFlags = unix.MAP_PRIVATE | unix.MAP_ANON
)
