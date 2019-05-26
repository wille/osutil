# osutil

[![Travis](https://img.shields.io/travis/wille/osutil/master.svg)](https://travis-ci.org/wille/osutil)
[![GoDoc](https://godoc.org/github.com/wille/osutil?status.svg)](https://godoc.org/github.com/wille/osutil)
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
