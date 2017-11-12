package config

import (
	"testing"
	"github.com/SenseException/go-phersion/config"
	"os"
	"io/ioutil"
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
	assertNoError(err, t)
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
	dir := os.TempDir() + "/test_directory"

	config.Init(dir)

	stat, err := os.Stat(dir + "/config.json")
	assertNoError(err, t)

	if stat.IsDir() {
		t.Error("config.json is not a file")
	}

	os.RemoveAll(dir)
}

// Expected json config is created
func TestCreateJson(t *testing.T) {
	dir := os.TempDir() + "/test_directory"

	config.Init(dir)

	configJson, err := ioutil.ReadFile(dir + "/config.json")
	assertNoError(err, t)

	var expected string = `{"major":1,"minor":0,"patch":0,"label":"","identifier":0}`

	if string(configJson) != expected {
		t.Errorf("Expected that config %s is equal to %s", configJson, expected)
	}

	os.RemoveAll(dir)
}

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Error(err.Error())
	}
}
