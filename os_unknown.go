//+build !windows,!darwin,!linux,!freebsd,!openbsd,!dragonfly,!netbsd,!solaris

package osutil

import "runtime"

const (
	Name = runtime.GOOS
)

func GetVersion() string {
	return ""
}

func GetDisplay() string {
	return Name
}
