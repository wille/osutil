package osutil

import (
	"os"
	"os/exec"
	"strings"
)

var dists []distroInfo

type distroInfo struct {
	Display string
	// IDs matches the os-release "ID" field (the canonical machine-readable
	// identifier on modern systems).
	IDs []string
	// Search matches the detected name reported by lsb_release or os-release.
	Search []string
	// Files matches the existence of a distribution-specific release file.
	Files []string
}

func init() {
	dists = []distroInfo{
		{"Debian", []string{"debian"}, nil, nil},
		{"Ubuntu", []string{"ubuntu"}, nil, nil},
		{"openSUSE", []string{"opensuse", "opensuse-leap", "opensuse-tumbleweed"}, []string{"SUSE Linux", "openSUSE project"}, []string{"/etc/SuSE-release"}},
		{"Mint Linux", []string{"linuxmint"}, []string{"LinuxMint", "mint"}, nil},
		{"Gentoo", []string{"gentoo"}, nil, []string{"/etc/gentoo-release"}},
		{"Fedora", []string{"fedora"}, nil, []string{"/etc/fedora-release"}},
		{"CentOS", []string{"centos"}, nil, []string{"/etc/centos-release"}},
		{"Oracle Linux", []string{"ol"}, nil, []string{"/etc/oracle-release"}},
		{"Arch Linux", []string{"arch", "archarm"}, []string{"archlinux", "archarm", "arch"}, nil},
		{"Kali Linux", []string{"kali"}, []string{"kali", "debian kali linux"}, nil},
	}
}

func getLSB() (map[string]string, bool) {
	lsb := make(map[string]string)

	proc := exec.Command("lsb_release", "-irc")
	raw, err := proc.Output()

	if err != nil {
		return nil, false
	}

	for _, line := range strings.Split(string(raw), "\n") {
		pair := strings.SplitN(line, ":", 2)
		if len(pair) != 2 {
			continue
		}

		k := strings.Trim(pair[0], " \t")
		v := strings.Trim(pair[1], " \t")
		lsb[k] = v
	}

	return lsb, true
}

func getOSRelease() (map[string]string, bool) {
	raw, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return nil, false
	}

	return parseOSRelease(string(raw)), true
}

// resolveDistro finds the matching distribution from the detected name and the
// os-release ID.
func resolveDistro(detect, id string) distroInfo {
	dl := strings.ToLower(detect)
	idl := strings.ToLower(id)

	for _, k := range dists {
		if dl != "" && strings.ToLower(k.Display) == dl {
			return k
		}

		if idl != "" {
			for _, x := range k.IDs {
				if strings.ToLower(x) == idl {
					return k
				}
			}
		}

		if dl != "" {
			for _, search := range k.Search {
				if strings.ToLower(search) == dl {
					return k
				}
			}
		}

		for _, file := range k.Files {
			if _, err := os.Stat(file); err == nil {
				return k
			}
		}
	}

	return distroInfo{}
}

// suseCodename derives the openSUSE product line (e.g. "Leap", "Tumbleweed")
// from the os-release ID, which is used as the codename since openSUSE does not
// expose VERSION_CODENAME.
func suseCodename(id string) string {
	const prefix = "opensuse-"
	if !strings.HasPrefix(strings.ToLower(id), prefix) {
		return ""
	}

	suffix := id[len(prefix):]
	if suffix == "" {
		return ""
	}

	return strings.ToUpper(suffix[:1]) + suffix[1:]
}

// dist resolves the distribution from already-gathered lsb_release and
// os-release data. lsb takes precedence; os-release fills in the gaps. Either
// map may be nil.
func dist(lsb, osmap map[string]string) Distro {
	detect := lsb["Distributor ID"]
	release := lsb["Release"]
	codename := lsb["Codename"]
	id := osmap["ID"]

	if detect == "" {
		if v := osmap["DISTRIB_ID"]; v != "" {
			detect = v
		} else if v := osmap["NAME"]; v != "" {
			detect = v
		}
	}

	if release == "" {
		if v := osmap["VERSION_ID"]; v != "" {
			release = v
		}
		if v := osmap["DISTRIB_RELEASE"]; v != "" {
			release = v
		}
	}

	if codename == "" {
		if v := osmap["VERSION_CODENAME"]; v != "" {
			codename = v
		} else if v := osmap["DISTRIB_CODENAME"]; v != "" {
			codename = v
		}
	}

	info := resolveDistro(detect, id)

	if codename == "" {
		codename = suseCodename(id)
	}

	return Distro{info.Display, release, codename}
}

func GetDist() Distro {
	lsb, _ := getLSB()
	osmap, _ := getOSRelease()
	return dist(lsb, osmap)
}
