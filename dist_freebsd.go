package osutil

import (
	"io/ioutil"
	"strings"
)

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
