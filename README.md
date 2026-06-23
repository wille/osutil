# osutil

[![CI](https://github.com/wille/osutil/actions/workflows/ci.yml/badge.svg)](https://github.com/wille/osutil/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/wille/osutil.svg)](https://pkg.go.dev/github.com/wille/osutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/wille/osutil)](https://goreportcard.com/report/github.com/wille/osutil)

Go library to detect operating system version

## Usage

```go
import "github.com/wille/osutil"

// macOS, Windows, Linux, FreeBSD...
osutil.Name

// macOS: 10.17, 10.16...
// Windows: 10, 8.1, Vista...
//
osutil.GetVersion()

// Returns nicely formatted string
// Windows 10, macOS 10.17 Sierra
osutil.GetDisplay()
```
