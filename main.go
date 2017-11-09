package main

import (
	"flag"
	"github.com/SenseException/go-phersion/config"
)

func main() {
	// Define arguments
	dirPath := flag.String("config", "./.go-pherson", "Path to version config file directory")
	flag.Parse()

	// Retrieve command
	//command := flag.Arg(0)

	var configPath string = *dirPath

	if ! config.Exists(configPath) {
		config.Init(configPath)
	}
}
