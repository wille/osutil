//go:build windows

package osutil

import "testing"

func TestEdition(t *testing.T) {
	cases := map[string]string{
		"10.0.18362.657": "10",
		"6.3.9600":       "8.1",
		"6.2.9200":       "8",
		"6.1.7601":       "7",
		"6.0.6002":       "Vista",
		"5.2.3790":       "Server 2003",
		"5.1.2600":       "XP",
		"5.0.2195":       "2000",
		"4.0.1381":       "", // unknown major.minor
		"10":             "", // too few components
		"":               "", // empty
	}

	for in, want := range cases {
		if got := edition(in); got != want {
			t.Errorf("edition(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestVersionRegexp guards the issue #3 fix: the version number must be
// extracted regardless of the localized word in the "ver" banner.
func TestVersionRegexp(t *testing.T) {
	cases := map[string]string{
		"Microsoft Windows [Version 10.0.18362.657]": "10.0.18362.657",
		"Microsoft Windows [版本 10.0.18362.657]":      "10.0.18362.657",
		"Microsoft Windows XP [Version 5.1.2600]":    "5.1.2600",
		"no version here": "",
	}

	for in, want := range cases {
		if got := versionRegexp.FindString(in); got != want {
			t.Errorf("versionRegexp.FindString(%q) = %q, want %q", in, got, want)
		}
	}
}
