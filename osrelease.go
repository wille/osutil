package osutil

import "strings"

// parseOSRelease parses the key=value contents of an os-release file (see
// https://www.freedesktop.org/software/systemd/man/os-release.html). Comments,
// blank lines, and malformed lines are skipped, and surrounding quotes are
// stripped from values.
func parseOSRelease(s string) map[string]string {
	osmap := make(map[string]string)

	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			continue
		}

		k := strings.TrimSpace(pair[0])
		v := strings.Trim(strings.TrimSpace(pair[1]), "\"")
		osmap[k] = v
	}

	return osmap
}
