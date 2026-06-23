package osutil

import (
	"os"
	"os/exec"
	"strings"
)

func getFreeBSDVersion() (string, bool) {
	proc := exec.Command("freebsd-version", "-u")
	raw, err := proc.Output()

	if err != nil {
		return "", false
	}

	return strings.TrimRight(string(raw), "\n"), true
}

func getOSRelease() (map[string]string, bool) {
	raw, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return nil, false
	}

	return parseOSRelease(string(raw)), true
}

func GetDist() Distro {
	var detect string
	var release string

	version, freeBSDVersionExists := getFreeBSDVersion()
	if freeBSDVersionExists {
		return Distro{"FreeBSD", version, ""}
	}

	osmap, osMapExists := getOSRelease()

	if osMapExists {
		name := osmap["NAME"]
		if name != "" {
			detect = name
		}

		version := osmap["VERSION_ID"]
		if version != "" {
			release = version
		}
	}

	return Distro{detect, release, ""}
}
