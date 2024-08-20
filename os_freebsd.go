package osutil

import "fmt"

const Name = "FreeBSD"

func GetDisplay() string {
	distro := GetDist()

	if distro.Display != "" {
		return fmt.Sprintf("%s %s", distro.Display, distro.Release)
	}

	return Name
}

func GetVersion() string {
	return ""
}
