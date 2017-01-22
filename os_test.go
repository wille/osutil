package oslib

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("Name:", Name)
	fmt.Println("Version:", GetVersion())
	fmt.Println("Display:", GetDisplay())
	fmt.Println("Arch:", GetDisplayArch())
}
