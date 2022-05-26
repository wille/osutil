package osutil

import "testing"

func TestWindows_parseVersion(t *testing.T) {
	for _, kv := range [][2]string{
		{"Microsoft Windows [版本 10.0.19044.1706]", "10.0.19044.1706"},
		{"Microsoft Windows [Version 10.0.19044.1706]", "10.0.19044.1706"},
		{"hello world[asdasd 01.121.1212]", "01.121.1212"},
		{"hello world[asdasd01.121.1212]", "01.121.1212"},
		{"hello world[asdasd01.121.1212 ]  12", "01.121.1212"},
	} {
		ver, err := parseVersion(kv[0])
		if ver != kv[1] {
			t.Logf("parse '%s' expected:'%s' actual:'%s' reason:%v", kv[0], kv[1], ver, err)
			t.FailNow()
		}
	}
}
