package versioning

import (
	"testing"
	"github.com/SenseException/go-phersion/versioning"
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
	expectedVersion(versioning.Version{Major: 2, Minor: 1, Patch: 9, Label: "beta", Identifier: 4}, "2.1.9-beta4", t)
}

func TestGetWithoutLabel(t *testing.T) {
	expectedVersion(versioning.Version{Major: 2, Minor: 1, Patch: 9, Identifier: 4}, "2.1.9", t)
}

func expectedVersion(version versioning.Version, expected string, t *testing.T) {
	versionStr := version.Get()

	if expected != versionStr {
		t.Errorf("Expected version was %s, but got %s", expected, versionStr)
	}
}

func TestAddType(t *testing.T) {
	version := versioning.Version{}
	version.AddType("foo")

	expected := []string{"foo"}

	if len(expected) != len(version.VersionTypes) {
		t.Error("Version type has not the expected length")
	}

	if expected[0] != version.VersionTypes[0] {
		t.Error("Version type foo was not added")
	}
}

func TestRemoveType(t *testing.T) {
	version := versioning.Version{Major: 1, Minor: 0, Patch: 0, Label: "", Identifier: 0, VersionTypes: []string{"foo", "bar", "baz"}}
	version.RemoveType("bar")

	expected := []string{"foo", "baz"}

	if len(expected) != len(version.VersionTypes) {
		t.Error("Version type has not the expected length")
	}

	if expected[0] != version.VersionTypes[0] || expected[1] != version.VersionTypes[1] {
		t.Error("Version type foo was not removed properly")
	}
}