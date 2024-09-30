package osutil

import (
	"io/ioutil"
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
	osmap := make(map[string]string)

	raw, err := ioutil.ReadFile("/etc/os-release")

	if err != nil {
		return osmap, false
	}

	s := string(raw)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			break
		}

		pair := strings.Split(line, "=")
		k := pair[0]
		v := pair[1]

		v = strings.Trim(v, "\"")

		osmap[k] = v
	}

	return osmap, true
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
