package osutil

import (
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

const (
	Name = "Windows"
)

// versionRegexp extracts the dotted version number (e.g. 10.0.18362.657) from
// the output of "ver". Matching the number directly avoids relying on the word
// "Version", which is localized on non-English Windows (e.g. "[版本 ...]") and
// previously caused a panic (issue #3).
var versionRegexp = regexp.MustCompile(`\d+(\.\d+)+`)

// getVersionNumber returns the dotted version number reported by "ver" (e.g.
// "10.0.18362.657"), or an empty string if it cannot be determined.
func getVersionNumber() string {
	cmd := exec.Command("cmd", "/c", "ver")
	// Hide the console window; that is the expected behavior in most cases.
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return versionRegexp.FindString(string(out))
}

// edition maps a dotted Windows version number to its marketing edition (e.g.
// "10.0.x" -> "10"). It returns an empty string for unknown or malformed input.
func edition(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return ""
	}

	switch parts[0] + "." + parts[1] {
	case "10.0": // 10 / 11 / Server 2016+
		return "10"
	case "6.3": // Server 2012 R2
		return "8.1"
	case "6.2": // Server 2012
		return "8"
	case "6.1":
		return "7"
	case "6.0":
		return "Vista"
	case "5.2":
		return "Server 2003"
	case "5.1":
		return "XP"
	case "5.0":
		return "2000"
	}

	return ""
}

func GetVersion() string {
	return edition(getVersionNumber())
}

func GetDisplay() string {
	display := Name
	version := GetVersion()

	if version != "" {
		display += " " + version
	}

	return display
}
