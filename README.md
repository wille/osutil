# osutil

[![Release](https://img.shields.io/github/v/release/wille/osutil)](https://github.com/wille/osutil/releases)
[![CI](https://github.com/wille/osutil/actions/workflows/ci.yml/badge.svg)](https://github.com/wille/osutil/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/wille/osutil.svg)](https://pkg.go.dev/github.com/wille/osutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/wille/osutil)](https://goreportcard.com/report/github.com/wille/osutil)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wille/osutil?logo=go)](https://go.dev)

Go library to detect the operating system, its version, distribution and architecture.

## Install

```sh
go get github.com/wille/osutil
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/wille/osutil"
)

func main() {
	fmt.Println(osutil.Name)             // e.g. "Linux"
	fmt.Println(osutil.GetVersion())     // e.g. "6.8.0-40-generic"
	fmt.Println(osutil.GetDisplay())     // e.g. "Ubuntu 24.04"
	fmt.Println(osutil.GetDisplayArch()) // e.g. "amd64"

	dist := osutil.GetDist()
	fmt.Println(dist.Display)  // e.g. "Ubuntu"
	fmt.Println(dist.Release)  // e.g. "24.04"
	fmt.Println(dist.Codename) // e.g. "noble"
}
```

## API

### `osutil.Name`

A constant holding the name of the operating system the binary was built for. It
equals one of the exported name constants:

```go
osutil.Windows      // "Windows"
osutil.MacOS        // "macOS"
osutil.Linux        // "Linux"
osutil.FreeBSD      // "FreeBSD"
osutil.OpenBSD      // "OpenBSD"
osutil.DragonflyBSD // "DragonFlyBSD"
osutil.NetBSD       // "NetBSD"
osutil.Solaris      // "Solaris"
```

On any other platform `Name` falls back to Go's `runtime.GOOS`.

### `GetVersion() string`

Returns the OS version. Its meaning is platform-specific:

| Platform | Returns | Example |
| --- | --- | --- |
| macOS | macOS release number | `15.0`, `10.15` |
| Windows | Windows edition | `10`, `8.1`, `XP` |
| Linux | Kernel version (`uname -r`) | `6.8.0-40-generic` |
| FreeBSD and others | empty string | `""` |

### `GetDisplay() string`

Returns a human-readable name combining the OS and version. Never empty.

| Platform | Example |
| --- | --- |
| macOS | `macOS 15.0 Sequoia` |
| Windows | `Windows 10` |
| Linux | `Ubuntu 24.04` (or `Linux` if the distro is unknown) |
| FreeBSD | `FreeBSD 13.2` (or `FreeBSD`) |

### `GetDist() Distro`

Returns distribution details as a `Distro` value:

```go
type Distro struct {
	Display  string // e.g. "Ubuntu", "openSUSE", "macOS"
	Release  string // e.g. "24.04", "15.3", "15.0"
	Codename string // e.g. "noble", "Leap", "Sequoia"
}
```

- **Linux** detects the distribution via `lsb_release` and `/etc/os-release`
  (Debian, Ubuntu, openSUSE, Mint, Gentoo, Fedora, CentOS, Oracle, Arch, Kali, ...).
- **macOS** maps the kernel version to the release (e.g. `{macOS, 15.0, Sequoia}`).
- **FreeBSD** reports the running version.
- On other platforms it returns the zero `Distro{}`.

### `GetDisplayArch() string`

Returns the architecture the binary was built for (`runtime.GOARCH`), e.g.
`amd64`, `arm64`, `386`.
