package osutil

import (
	"runtime"
	"testing"
)

// TestNameMatchesGOOS asserts the Name constant matches the OS we are actually
// running on.
func TestNameMatchesGOOS(t *testing.T) {
	want := map[string]string{
		"windows":   Windows,
		"darwin":    MacOS,
		"linux":     Linux,
		"freebsd":   FreeBSD,
		"openbsd":   OpenBSD,
		"dragonfly": DragonflyBSD,
		"netbsd":    NetBSD,
		"solaris":   Solaris,
	}

	if exp, ok := want[runtime.GOOS]; ok {
		if Name != exp {
			t.Errorf("Name = %q, want %q for GOOS=%q", Name, exp, runtime.GOOS)
		}
	} else if Name != runtime.GOOS {
		// Unrecognized platforms fall back to runtime.GOOS (os_unknown.go).
		t.Errorf("Name = %q, want %q (runtime.GOOS)", Name, runtime.GOOS)
	}
}

func TestDisplayArch(t *testing.T) {
	if got := GetDisplayArch(); got != runtime.GOARCH {
		t.Errorf("GetDisplayArch() = %q, want %q", got, runtime.GOARCH)
	}
}

func TestGetDisplayNonEmpty(t *testing.T) {
	if GetDisplay() == "" {
		t.Error("GetDisplay() returned an empty string")
	}
}

// TestLogDetected has no assertions; it records what was detected so the values
// are visible in CI logs across all three runner OSes.
func TestLogDetected(t *testing.T) {
	t.Logf("Name=%q Version=%q Display=%q Arch=%q",
		Name, GetVersion(), GetDisplay(), GetDisplayArch())
	t.Logf("Dist=%+v", GetDist())
}
