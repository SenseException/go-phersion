package file

import (
	"os"
	"strings"
)

func Init(filePath string) {
	os.MkdirAll(filePath, 0644)
}

func Exists(dirPath string) bool {
	dirInfo, dirNotFound := os.Stat(dirPath)
	fileInfo, fileNotFound := os.Stat(getFilePath(dirPath))

	dirExists := os.IsExist(dirNotFound) || dirInfo.IsDir()
	fileExists := os.IsExist(fileNotFound) || ! fileInfo.IsDir()

	return dirExists && fileExists
}

func getFilePath(dirPath string) string {
	dirPath = strings.TrimRight(dirPath, os.PathSeparator)

	return dirPath + os.PathSeparator + "config.json"
}