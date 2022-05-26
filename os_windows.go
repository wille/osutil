package osutil

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

const (
	Name = "Windows"
)

func getEdition() string {
	cmd := exec.Command("cmd")

	var out bytes.Buffer
	cmd.Stdout = &out
	// the next 2 lines makes a console window hidden, that is expected behavior in most cases.
	cmd.SysProcAttr = new(syscall.SysProcAttr)
	cmd.SysProcAttr.HideWindow = true

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	raw := out.String()
	ver, err := parseVersion(raw)
	if err != nil {
		panic(err)
	}
	return ver
}

func GetVersion() string {
	version := getEdition()
	if version == "" {
		return ""
	}
	parts := strings.Split(version, ".")
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

func GetDist() Distro {
	return Distro{
		Display:  Name,
		Release:  getEdition(),
		Codename: "",
	}
}

func parseVersion(raw string) (string, error) {
	// NOTE: state
	//  0-start
	//  1-inside
	//  2-version found
	//  3-stop
	var state = 0
	newErr := func(reason string) error {
		return fmt.Errorf("invalid version:%s: reason:%s", raw, reason)
	}
	// first and last version character
	var i1, i2 = -1, -1
	for i := 0; i < len(raw); i++ {
		if state == 3 {
			break
		}
		switch raw[i] {
		case '[':
			if state != 0 {
				return "", newErr("too many '['")
			}
			state = 1
		case ']':
			if state <= 0 {
				return "", newErr("']' must bebind '['")
			}
			state = 3
		case ' ':
			// pass
		default:
			// isdigit or isdot
			if s := raw[i]; (s >= '0' && s <= '9') || s == '.' {
				if state == 1 {
					i1 = i // first char
					state = 2
				}
				i2 = i // last char
			}
		}
	}
	if i1 == -1 || i2 == -1 {
		return "", newErr("missing index")
	}
	return raw[i1 : i2+1], nil
}
