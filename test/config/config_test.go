package config

import (
	"testing"
	"github.com/SenseException/go-phersion/config"
	"os"
)

/*
 Tests for the method config.Exists
*/

// Return false, because config directory doesn't exists
func TestDirDoesNotExist(t *testing.T) {
	dir := os.TempDir() + "/test_directory"

	if config.Exists(dir) {
		t.Errorf("%s was not expected to be found", dir)
	}
}

// Return false, because config file doesn't exists
func TestConfigDoesNotExist(t *testing.T) {
	dir := os.TempDir() + "/test_directory"
	os.Mkdir(dir, 0777)

	if config.Exists(dir) {
		t.Error("config.json was not expected to be found")
	}

	os.RemoveAll(dir)
}

// Return true, because config directory and file exists
func TestConfigExist(t *testing.T) {
	dir := os.TempDir() + "/test_directory"
	os.Mkdir(dir, 0777)

	f, err := os.Create(dir + "/config.json")

	if err != nil {
		t.Error(err.Error())
	}
	f.Close()

	if false == config.Exists(dir) {
		t.Errorf("A config directory structure was expected in %s", dir)
	}

	os.RemoveAll(dir)
}

/*
 Tests for initializing config structure
*/

// Config structure is fully initialized
func TestCreateConfig(t *testing.T) {
	//os.Stdin

	dir := os.TempDir() + "/test_directory"

	config.Init(dir)

	stat, err := os.Stat(dir + "/config.json")
	if err != nil {
		t.Error(err.Error())
	}

	if stat.IsDir() {
		t.Error("config.json is not a file")
	}

	os.RemoveAll(dir)
}
