//go:build linux

package osutil

import "testing"

// osReleaseOpenSUSE is the /etc/os-release reported in issue #7 (openSUSE Leap
// 15.3), which has no lsb_release and no /etc/SuSE-release.
const osReleaseOpenSUSE = `NAME="openSUSE Leap"
VERSION="15.3"
ID="opensuse-leap"
ID_LIKE="suse opensuse"
VERSION_ID="15.3"
PRETTY_NAME="openSUSE Leap 15.3"
ANSI_COLOR="0;32"
CPE_NAME="cpe:/o:opensuse:leap:15.3"
BUG_REPORT_URL="https://bugs.opensuse.org"
HOME_URL="https://www.opensuse.org/"
`

func TestParseOSReleaseOpenSUSE(t *testing.T) {
	m := parseOSRelease(osReleaseOpenSUSE)

	for k, want := range map[string]string{
		"NAME":       "openSUSE Leap",
		"ID":         "opensuse-leap",
		"VERSION_ID": "15.3",
	} {
		if got := m[k]; got != want {
			t.Errorf("%s = %q, want %q", k, got, want)
		}
	}
}

func TestResolveDistroOpenSUSE(t *testing.T) {
	m := parseOSRelease(osReleaseOpenSUSE)

	if info := resolveDistro(m["NAME"], m["ID"]); info.Display != "openSUSE" {
		t.Errorf("Display = %q, want openSUSE", info.Display)
	}
}

func TestSuseCodename(t *testing.T) {
	for id, want := range map[string]string{
		"opensuse-leap":       "Leap",
		"opensuse-tumbleweed": "Tumbleweed",
		"ubuntu":              "",
		"":                    "",
	} {
		if got := suseCodename(id); got != want {
			t.Errorf("suseCodename(%q) = %q, want %q", id, got, want)
		}
	}
}

func TestParseOSReleaseSkipsCommentsAndBlanks(t *testing.T) {
	const content = `# a comment

ID=fedora
VERSION_ID=39
malformed line without equals
`
	m := parseOSRelease(content)

	if m["ID"] != "fedora" {
		t.Errorf("ID = %q, want fedora", m["ID"])
	}
	if m["VERSION_ID"] != "39" {
		t.Errorf("VERSION_ID = %q, want 39", m["VERSION_ID"])
	}
}
