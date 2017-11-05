package config

import (
	"os"
	"strings"
	"path/filepath"
)

func Init(filePath string) {
	os.MkdirAll(filePath, 0644)
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