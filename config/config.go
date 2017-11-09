package config

import (
	"os"
	"strings"
	"path/filepath"
	"fmt"
)

func Init(filePath string) {
	isInit := scanBool("Initializing Go-Phersion version config? [Y/n]", "Y")

	if (isInit) {
		os.MkdirAll(filePath, 0744)
		f, err := os.Create(filePath + "/config.json")

		if err != nil {
			fmt.Println(err.Error())
		}
		f.Close()
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