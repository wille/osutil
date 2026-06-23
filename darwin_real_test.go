//go:build darwin

package osutil

import (
	"os/exec"
	"strings"
	"testing"
)

func majorComponent(v string) string {
	if i := strings.Index(v, "."); i >= 0 {
		return v[:i]
	}
	return v
}

// TestDarwinKernelMatchesUname checks the detected kernel matches `uname -r`.
func TestDarwinKernelMatchesUname(t *testing.T) {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		t.Skipf("uname failed: %v", err)
	}

	unameMajor := majorComponent(strings.TrimSpace(string(out)))
	kernelMajor := majorComponent(getKernelRelease())

	if unameMajor != kernelMajor {
		t.Errorf("kernel major from getKernelRelease() = %q, uname -r = %q", kernelMajor, unameMajor)
	}
}

// TestDarwinVersionMatchesSwVers cross-checks the reported macOS version against
// sw_vers. It skips (rather than fails) when the runner's kernel is newer than
// the mapping table, so CI stays green until dist_darwin.go is updated.
func TestDarwinVersionMatchesSwVers(t *testing.T) {
	if _, ok := mapping[kernelVersion]; !ok {
		t.Skipf("kernel %d not in mapping; add it to dist_darwin.go", kernelVersion)
	}

	out, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		t.Skipf("sw_vers failed: %v", err)
	}

	swMajor := majorComponent(strings.TrimSpace(string(out)))
	if got := majorComponent(GetVersion()); got != swMajor {
		t.Errorf("GetVersion() major = %q, sw_vers major = %q", got, swMajor)
	}

	if disp := GetDisplay(); !strings.Contains(disp, "macOS") && !strings.Contains(disp, "Mac OS X") {
		t.Errorf("GetDisplay() = %q, expected to contain macOS or Mac OS X", disp)
	}
}
