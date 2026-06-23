package osutil

import (
	"os"
	"strings"
)

const (
	Name = "Linux"
)

func GetDisplay() string {
	distro := GetDist()

	if distro.Display != "" {
		return distro.Display + " " + distro.Release
	}

	return Name
}

// GetVersion returns the running kernel version (the equivalent of "uname -r",
// e.g. "6.18.6-arch1-1"), or an empty string if it cannot be determined.
func GetVersion() string {
	raw, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(raw))
}
