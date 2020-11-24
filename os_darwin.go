package osutil

const (
	Name = "macOS"
)

func GetVersion() string {
	if v, ok := mapping[kernelVersion]; ok {
		return v.release
	}

	return ""
}

func GetDisplay() string {
	if release, ok := mapping[kernelVersion]; ok {
		return release.name + " " + release.release + " " + release.releaseName
	}

	return "macOS"
}
