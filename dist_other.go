//go:build !linux && !darwin && !freebsd

package osutil

// GetDist returns an empty Distro on platforms without distribution
// information (Windows, the remaining BSDs, Solaris, ...). It exists so the
// API is uniform across all platforms.
func GetDist() Distro {
	return Distro{}
}
