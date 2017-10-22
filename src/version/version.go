package version

import (
	"bytes"
	"strconv"
)

type Version struct {
	Major      int
	Minor      int
	Patch      int
	Label      string
	Identifier int
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
			versionBuffer.WriteString(".")
			versionBuffer.WriteString(strconv.Itoa(v.Identifier))
		}
	}

	return versionBuffer.String()
}