package osutil

const (
	Name = "macOS"
)

func GetVersion() string {
	return versionForKernel(kernelVersion)
}

func GetDisplay() string {
	return displayForKernel(kernelVersion)
}
