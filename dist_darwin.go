package osutil

import (
	"strconv"
	"strings"
)

type version struct {
	release, name, releaseName string
}

// mapping maps the Darwin kernel major version to the macOS release it ships
// with. Updated from https://en.wikipedia.org/wiki/MacOS_version_history#Releases
var mapping = map[int]version{
	25: {"26.0", "macOS", "Tahoe"},
	24: {"15.0", "macOS", "Sequoia"},
	23: {"14.0", "macOS", "Sonoma"},
	22: {"13.0", "macOS", "Ventura"},
	21: {"12.0", "macOS", "Monterey"},
	20: {"11.0", "macOS", "Big Sur"},
	19: {"10.15", "macOS", "Catalina"},
	18: {"10.14", "macOS", "Mojave"},
	17: {"10.13", "macOS", "High Sierra"},
	16: {"10.12", "macOS", "Sierra"},
	15: {"10.11", "Mac OS X", "El Capitan"},
	14: {"10.10", "Mac OS X", "Yosemite"},
	13: {"10.9", "Mac OS X", "Mavericks"},
	12: {"10.8", "Mac OS X", "Mountain Lion"},
	11: {"10.7", "Mac OS X", "Lion"},
	10: {"10.6", "Mac OS X", "Snow Leopard"},
	9:  {"10.5", "Mac OS X", "Leopard"},
	8:  {"10.4", "Mac OS X", "Tiger"},
	7:  {"10.3", "Mac OS X", "Panther"},
	6:  {"10.2", "Mac OS X", "Jaguar"},
	5:  {"10.1", "Mac OS X", "Puma"},
}

var kernelVersion int

func init() {
	kernelVersion = parseKernelMajor(getKernelRelease())
}

// parseKernelMajor extracts the leading integer from a Darwin kernel release
// string (e.g. "23.6.0" -> 23). It returns 0 if the string cannot be parsed.
func parseKernelMajor(release string) int {
	i := strings.Index(release, ".")
	if i < 0 {
		i = len(release)
	}

	n, _ := strconv.Atoi(release[:i])
	return n
}

func distForKernel(k int) Distro {
	if v, ok := mapping[k]; ok {
		return Distro{v.name, v.release, v.releaseName}
	}

	return Distro{}
}

func versionForKernel(k int) string {
	if v, ok := mapping[k]; ok {
		return v.release
	}

	return ""
}

func displayForKernel(k int) string {
	if v, ok := mapping[k]; ok {
		return v.name + " " + v.release + " " + v.releaseName
	}

	return "macOS"
}

func GetDist() Distro {
	return distForKernel(kernelVersion)
}
