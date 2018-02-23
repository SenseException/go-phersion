package versioning

import (
	"bytes"
	"strconv"
)

type Version struct {
	Major        int
	Minor        int
	Patch        int
	Label        string
	Identifier   int
	VersionTypes []string
}

func (v *Version) Get() string {
	var versionBuffer bytes.Buffer

	versionBuffer.WriteString(strconv.Itoa(v.Major))
	versionBuffer.WriteString(".")
	versionBuffer.WriteString(strconv.Itoa(v.Minor))
	versionBuffer.WriteString(".")
	versionBuffer.WriteString(strconv.Itoa(v.Patch))

	if "" != v.Label {
		versionBuffer.WriteString("-")
		versionBuffer.WriteString(v.Label)

		if 0 != v.Identifier {
			versionBuffer.WriteString(strconv.Itoa(v.Identifier))
		}
	}

	return versionBuffer.String()
}

func (v *Version) AddType(versionType string) {
	// TODO prevent duplications of version types
	v.VersionTypes = append(v.VersionTypes, versionType)
}

func (v *Version) RemoveType(versionType string) {
	for i := 0; i < len(v.VersionTypes); i++  {
		if v.VersionTypes[i] == versionType {
			v.VersionTypes = append(v.VersionTypes[:i], v.VersionTypes[i+1:]...)
			break
		}
	}
}