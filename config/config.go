package config

import (
	"os"
	"strings"
	"path/filepath"
	"fmt"
	"github.com/SenseException/go-phersion/versioning"
	"io/ioutil"
	"encoding/json"
)

func Write(version versioning.Version, dirPath string) error {
	jsonConfig, _ := json.Marshal(createConfig(version))
	fmt.Println(string(jsonConfig))

	err := ioutil.WriteFile(getFilePath(dirPath), jsonConfig, 0744)

	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func Init(dirPath string) error {
	isInit := scanBool("Initializing Go-Phersion version config? [Y/n]", "Y")
	var err error = nil

	if (isInit) {
		os.MkdirAll(dirPath, 0744)
		err = Write(versioning.Version{Major: 1}, dirPath)
	}

	return err
}

func Exists(dirPath string) bool {
	dirInfo, dirNotFound := os.Stat(dirPath)
	fileInfo, fileNotFound := os.Stat(getFilePath(dirPath))

	dirExists := nil == dirNotFound && dirInfo.IsDir()
	fileExists := nil == fileNotFound && ! fileInfo.IsDir()

	return dirExists && fileExists
}

func getFilePath(dirPath string) string {
	dirPath = strings.TrimRight(dirPath, string(os.PathSeparator))

	return filepath.FromSlash(dirPath + "/config.json")
}

func scanBool(question string, preset string) bool {
	fmt.Println(question)
	fmt.Scanf("%s", &preset)

	return "Y" == preset || "y" == preset
}

func createConfig(version versioning.Version) configJson {
	return configJson{
		Major: version.Major,
		Minor: version.Minor,
		Patch: version.Patch,
		Label: version.Label,
		Identifier: version.Identifier,
	}
}

type configJson struct {
	Major      int		`json:"major"`
	Minor      int		`json:"minor"`
	Patch      int		`json:"patch"`
	Label      string	`json:"label"`
	Identifier int		`json:"identifier"`
}