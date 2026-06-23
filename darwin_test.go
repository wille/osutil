//go:build darwin

package osutil

import "testing"

func TestParseKernelMajor(t *testing.T) {
	cases := map[string]int{
		"23.6.0": 23,
		"19.6.0": 19,
		"24":     24, // no dot
		"":       0,  // empty
		"x.y":    0,  // non-numeric
	}

	for in, want := range cases {
		if got := parseKernelMajor(in); got != want {
			t.Errorf("parseKernelMajor(%q) = %d, want %d", in, got, want)
		}
	}
}

func TestVersionForKernel(t *testing.T) {
	if got := versionForKernel(19); got != "10.15" {
		t.Errorf("versionForKernel(19) = %q, want 10.15", got)
	}
	if got := versionForKernel(999); got != "" {
		t.Errorf("versionForKernel(999) = %q, want empty", got)
	}
}

func TestDisplayForKernel(t *testing.T) {
	if got := displayForKernel(19); got != "macOS 10.15 Catalina" {
		t.Errorf("displayForKernel(19) = %q, want 'macOS 10.15 Catalina'", got)
	}
	// Unknown kernel falls back to the bare OS name.
	if got := displayForKernel(999); got != "macOS" {
		t.Errorf("displayForKernel(999) = %q, want 'macOS'", got)
	}
}

func TestDistForKernel(t *testing.T) {
	want := Distro{"macOS", "10.15", "Catalina"}
	if got := distForKernel(19); got != want {
		t.Errorf("distForKernel(19) = %+v, want %+v", got, want)
	}
	if got := distForKernel(999); got != (Distro{}) {
		t.Errorf("distForKernel(999) = %+v, want empty Distro", got)
	}
}
