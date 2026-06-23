package osutil

const (
	Name = "Linux"
)

func GetDisplay() string {
	distro := GetDist()

	if distro.Display != "" {
		return distro.Display + " " + distro.Release
	}

	return Name
}

// GetVersion returns the distribution release (e.g. "20.04", "15.3"), or an
// empty string if it cannot be determined.
func GetVersion() string {
	return GetDist().Release
}
