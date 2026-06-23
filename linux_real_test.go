//go:build linux

package osutil

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

// TestLinuxKernelVersion asserts GetVersion returns the running kernel version.
func TestLinuxKernelVersion(t *testing.T) {
	raw, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		t.Skipf("cannot read /proc/sys/kernel/osrelease: %v", err)
	}

	want := strings.TrimSpace(string(raw))
	if got := GetVersion(); got != want {
		t.Errorf("GetVersion() = %q, want %q", got, want)
	}

	if !regexp.MustCompile(`^\d+\.\d+`).MatchString(want) {
		t.Errorf("kernel version %q does not look like X.Y...", want)
	}
}

// TestLinuxDistDetected cross-checks distro detection against /etc/os-release.
// On GitHub's ubuntu runner this exercises the "ubuntu" ID path end-to-end.
func TestLinuxDistDetected(t *testing.T) {
	raw, err := os.ReadFile("/etc/os-release")
	if err != nil {
		t.Skipf("no /etc/os-release: %v", err)
	}

	osmap := parseOSRelease(string(raw))
	id := osmap["ID"]

	d := GetDist()
	t.Logf("GetDist() = %+v (os-release ID=%q)", d, id)

	// Only assert when we actually recognize the running distribution.
	if resolveDistro("", id).Display == "" {
		t.Skipf("distro ID %q not in the known list", id)
	}

	if d.Display == "" {
		t.Errorf("GetDist().Display empty for known ID %q", id)
	}
	if !strings.Contains(GetDisplay(), d.Display) {
		t.Errorf("GetDisplay() = %q does not contain %q", GetDisplay(), d.Display)
	}
	// Fixed-release distros expose VERSION_ID; rolling ones (e.g. Arch) may not.
	if osmap["VERSION_ID"] != "" && d.Release == "" {
		t.Errorf("GetDist().Release empty despite VERSION_ID=%q", osmap["VERSION_ID"])
	}
}
