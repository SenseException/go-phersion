package file

import "github.com/SenseException/go-phersion/versioning"

type VersionFile interface {
	Format(version versioning.Version) string
}
