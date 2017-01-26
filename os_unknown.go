//+build !windows,!darwin,!linux

package oslib

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
