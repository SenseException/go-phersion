package versioning

import (
	"testing"
	"github.com/SenseException/go-phersion/src/versioning"
)

func TestGetWithMajor(t *testing.T) {
	expectedVersion(versioning.Version{Major: 42}, "42.0.0", t)
}

func TestGetWithMinor(t *testing.T) {
	expectedVersion(versioning.Version{Major: 2, Minor: 1}, "2.1.0", t)
}

func TestGetWithPatch(t *testing.T) {
	expectedVersion(versioning.Version{Major: 2, Minor: 1, Patch: 9}, "2.1.9", t)
}

func TestGetWithLabel(t *testing.T) {
	expectedVersion(versioning.Version{Major: 2, Minor: 1, Patch: 9, Label: "beta"}, "2.1.9-beta", t)
}

func TestGetWithIdentifier(t *testing.T) {
	expectedVersion(versioning.Version{Major: 2, Minor: 1, Patch: 9, Label: "beta", Identifier: 4}, "2.1.9-beta.4", t)
}

func expectedVersion(version versioning.Version, expected string, t *testing.T) {
	versionStr := version.Get()

	if expected != versionStr {
		t.Errorf("Expected version was %s, but got %s", expected, versionStr)
	}
}