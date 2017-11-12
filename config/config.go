package config

import (
	"os"
	"strings"
	"path/filepath"
	"fmt"
	"github.com/SenseException/go-phersion/versioning"
)

func Write(version versioning.Version, dirPath string) {
	config, err := os.Create(getFilePath(dirPath))

	if err != nil {
		fmt.Println(err.Error())
	}
	config.Close()
}

func Init(dirPath string) {
	isInit := scanBool("Initializing Go-Phersion version config? [Y/n]", "Y")

	if (isInit) {
		os.MkdirAll(dirPath, 0744)
		Write(versioning.Version{Major: 1}, dirPath)
	}
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