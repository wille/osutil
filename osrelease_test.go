package osutil

import "testing"

func TestParseOSRelease(t *testing.T) {
	cases := []struct {
		name    string
		content string
		want    map[string]string
	}{
		{
			name: "ubuntu",
			content: `NAME="Ubuntu"
VERSION="20.04.6 LTS (Focal Fossa)"
ID=ubuntu
ID_LIKE=debian
VERSION_ID="20.04"
VERSION_CODENAME=focal
`,
			want: map[string]string{
				"NAME":             "Ubuntu",
				"ID":               "ubuntu",
				"ID_LIKE":          "debian",
				"VERSION_ID":       "20.04",
				"VERSION_CODENAME": "focal",
			},
		},
		{
			name: "openSUSE Leap",
			content: `NAME="openSUSE Leap"
ID="opensuse-leap"
ID_LIKE="suse opensuse"
VERSION_ID="15.3"
PRETTY_NAME="openSUSE Leap 15.3"
`,
			want: map[string]string{
				"NAME":       "openSUSE Leap",
				"ID":         "opensuse-leap",
				"VERSION_ID": "15.3",
			},
		},
		{
			name: "comments blanks and malformed lines",
			content: `# this is a comment

ID=fedora
   VERSION_ID=39
malformed line without equals
EMPTY=
QUOTED="with spaces"
`,
			want: map[string]string{
				"ID":         "fedora",
				"VERSION_ID": "39",
				"EMPTY":      "",
				"QUOTED":     "with spaces",
			},
		},
		{
			name:    "empty input",
			content: "",
			want:    map[string]string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := parseOSRelease(tc.content)

			for k, want := range tc.want {
				if got[k] != want {
					t.Errorf("[%s] = %q, want %q", k, got[k], want)
				}
			}

			// "malformed line without equals" must not become a key.
			if _, ok := got["malformed line without equals"]; ok {
				t.Errorf("malformed line was parsed as a key: %v", got)
			}
		})
	}
}
