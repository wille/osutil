//go:build windows

package osutil

import (
	"strings"
	"testing"
)

func TestWindowsVersionNumber(t *testing.T) {
	v := getVersionNumber()
	if v == "" {
		t.Fatal("getVersionNumber() returned an empty string")
	}
	if !versionRegexp.MatchString(v) {
		t.Errorf("getVersionNumber() = %q is not a dotted version number", v)
	}
}

func TestWindowsDisplay(t *testing.T) {
	if disp := GetDisplay(); !strings.HasPrefix(disp, "Windows") {
		t.Errorf("GetDisplay() = %q, want prefix %q", disp, "Windows")
	}
}

func TestWindowsVersionNonEmpty(t *testing.T) {
	if GetVersion() == "" {
		t.Error("GetVersion() returned an empty string on Windows")
	}
}
