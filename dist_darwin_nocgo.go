//+build darwin,!cgo

package osutil

import (
	"log"
	"os/exec"
)

func getKernelRelease() string {
	out, err := exec.Command("sysctl", "-n", "kern.osrelease").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
