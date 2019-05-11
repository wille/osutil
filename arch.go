package osutil

import (
	"runtime"
)

func GetDisplayArch() string {
	arch := runtime.GOARCH

	return arch
}
