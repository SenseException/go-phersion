package config

import (
	"testing"
	"github.com/SenseException/go-phersion/config"
	"os"
)

func TestDirDoesNotExist(t *testing.T) {
	dir := os.TempDir() + "/test_directory"

	if config.Exists(dir) {
		t.Errorf("%s was not expected to be found", dir)
	}
}