package oslib

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var dists []distroInfo

type distroInfo struct {
	Display string
	Search  []string
	Files   []string
}

type Distro struct {
	Display  string
	Release  string
	Codename string
}

func init() {
	dists = make([]distroInfo, 3)
	dists[0] = distroInfo{"Debian", nil, nil}
	dists[1] = distroInfo{"Ubuntu", nil, nil}
	dists[2] = distroInfo{"openSUSE", []string{"SUSE Linux", "openSUSE project"}, []string{"/etc/SuSE-release"}}
}

func getLSB() (map[string]string, bool) {
	lsb := make(map[string]string)

	proc := exec.Command("lsb_release", "-irc")
	raw, _ := proc.Output()

	s := string(raw)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			break
		}

		pair := strings.Split(line, ":")
		k := pair[0]
		v := pair[1]

		k = strings.Trim(k, " \t")
		v = strings.Trim(v, " \t")

		lsb[k] = v
	}

	return lsb, true
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
	var codename string

	var info distroInfo

	lsb, lsbExists := getLSB()

	if lsbExists {
		detect = lsb["Distributor ID"]
		release = lsb["Release"]
		codename = lsb["Codename"]
	}

	osmap, osMapExists := getOSRelease()

	if !lsbExists && osMapExists {
		distribID := osmap["DISTRIB_ID"]
		if distribID != "" {
			detect = distribID
		}

		name := osmap["NAME"]
		if distribID == "" && name != "" {
			detect = name
		}

		version := osmap["VERSION_ID"]
		if version != "" {
			release = version
		}

		distribRelease := osmap["DISTRIB_RELEASE"]
		if distribRelease != "" {
			release = distribRelease
		}

		distribCodename := osmap["DISTRIB_CODENAME"]
		if distribCodename != "" {
			codename = distribCodename
		}
	}

	for _, k := range dists {
		dl := strings.ToLower(detect)

		if strings.ToLower(k.Display) == dl {
			info = k
			goto ret
		}

		if k.Search != nil {
			for _, search := range k.Search {
				if search == dl {
					info = k
					goto ret
				}
			}
		}

		if k.Files != nil {
			for _, file := range k.Files {
				if _, err := os.Stat(file); err == nil {
					info = k
					goto ret
				}
			}
		}
	}

ret:

	return Distro{info.Display, release, codename}
}
