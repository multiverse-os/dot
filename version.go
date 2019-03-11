package dot

import (
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (self Version) InvalidVersion() bool {
	return (self.Major == 0 && self.Minor == 0 && self.Patch == 0)
}

func (self Version) String() string {
	if self.InvalidVersion() {
		return "-"
	} else {
		return strconv.Itoa(self.Major) + "." + strconv.Itoa(self.Minor) + "." + strconv.Itoa(self.Patch)
	}
}

func MarshalVersion(version string) Version {
	versionComponents := strings.Split(version, ".")
	switch len(versionComponents) {
	case 1:
		major, _ := strconv.Atoi(versionComponents[0])
		return Version{Major: major}
	case 2:
		major, _ := strconv.Atoi(versionComponents[0])
		minor, _ := strconv.Atoi(versionComponents[1])
		return Version{Major: major, Minor: minor}
	case 3:
		major, _ := strconv.Atoi(versionComponents[0])
		minor, _ := strconv.Atoi(versionComponents[1])
		patch, _ := strconv.Atoi(versionComponents[2])
		return Version{Major: major, Minor: minor, Patch: patch}
	default:
		return Version{Major: 0, Minor: 0, Patch: 0}
	}
}
