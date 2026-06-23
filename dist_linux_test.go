//go:build linux

package osutil

import "testing"

func TestResolveDistro(t *testing.T) {
	cases := []struct {
		detect, id string
		want       string
	}{
		{"Ubuntu", "ubuntu", "Ubuntu"},
		{"Debian GNU/Linux", "debian", "Debian"},
		{"openSUSE Leap", "opensuse-leap", "openSUSE"},
		{"openSUSE Tumbleweed", "opensuse-tumbleweed", "openSUSE"},
		{"Fedora Linux", "fedora", "Fedora"},
		{"CentOS Linux", "centos", "CentOS"},
		{"Arch Linux", "arch", "Arch Linux"},
		{"Linux Mint", "linuxmint", "Mint Linux"},
		{"Oracle Linux Server", "ol", "Oracle Linux"},
		{"Kali GNU/Linux", "kali", "Kali Linux"},
		// Display-name match still works when ID is absent (lsb_release path).
		{"Ubuntu", "", "Ubuntu"},
		// Unknown distro resolves to an empty Display.
		{"Some New Distro", "newdistro", ""},
	}

	for _, tc := range cases {
		if got := resolveDistro(tc.detect, tc.id).Display; got != tc.want {
			t.Errorf("resolveDistro(%q, %q).Display = %q, want %q", tc.detect, tc.id, got, tc.want)
		}
	}
}

func TestSuseCodename(t *testing.T) {
	cases := map[string]string{
		"opensuse-leap":       "Leap",
		"opensuse-tumbleweed": "Tumbleweed",
		"OPENSUSE-LEAP":       "LEAP", // prefix match is case-insensitive, suffix is preserved
		"ubuntu":              "",
		"opensuse-":           "",
		"":                    "",
	}

	for id, want := range cases {
		if got := suseCodename(id); got != want {
			t.Errorf("suseCodename(%q) = %q, want %q", id, got, want)
		}
	}
}

func TestDist(t *testing.T) {
	cases := []struct {
		name string
		lsb  map[string]string
		os   map[string]string
		want Distro
	}{
		{
			name: "openSUSE Leap from os-release (issue #7)",
			lsb:  nil,
			os: map[string]string{
				"NAME":       "openSUSE Leap",
				"ID":         "opensuse-leap",
				"VERSION_ID": "15.3",
			},
			want: Distro{"openSUSE", "15.3", "Leap"},
		},
		{
			name: "Ubuntu from os-release with codename",
			lsb:  nil,
			os: map[string]string{
				"NAME":             "Ubuntu",
				"ID":               "ubuntu",
				"VERSION_ID":       "20.04",
				"VERSION_CODENAME": "focal",
			},
			want: Distro{"Ubuntu", "20.04", "focal"},
		},
		{
			name: "lsb_release takes precedence over os-release",
			lsb: map[string]string{
				"Distributor ID": "Debian",
				"Release":        "11",
				"Codename":       "bullseye",
			},
			os: map[string]string{
				"ID":               "debian",
				"VERSION_ID":       "10",
				"VERSION_CODENAME": "buster",
			},
			want: Distro{"Debian", "11", "bullseye"},
		},
		{
			name: "unknown distro keeps release but empty display",
			lsb:  nil,
			os: map[string]string{
				"NAME":       "Some New Distro",
				"ID":         "newdistro",
				"VERSION_ID": "1.0",
			},
			want: Distro{"", "1.0", ""},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := dist(tc.lsb, tc.os); got != tc.want {
				t.Errorf("dist() = %+v, want %+v", got, tc.want)
			}
		})
	}
}
