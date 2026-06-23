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

func getEdition() string {
	cmd := exec.Command("cmd", "/c", "ver")
	// Hide the console window; that is the expected behavior in most cases.
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return versionRegexp.FindString(string(out))
}

func GetVersion() string {
	version := getEdition()
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return ""
	}
	majormin := parts[0] + "." + parts[1]

	var edition string

	switch majormin {
	case "10.0": // 10 Server
		edition = "10"
	case "6.3": // Server 2012 R2
		edition = "8.1"
	case "6.2": // Server 2012
		edition = "8"
	case "6.1":
		edition = "7"
	case "6.0":
		edition = "Vista"
	case "5.2":
		edition = "Server 2003"
	case "5.1":
		edition = "XP"
	case "5.0":
		edition = "2000"
	}

	return edition
}

func GetDisplay() string {
	display := Name
	version := GetVersion()

	if version != "" {
		display += " " + version
	}

	return display
}
