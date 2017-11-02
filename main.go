package main

import (
	"fmt"
	"flag"
)

func main() {
	// Define arguments
	dirPath := flag.String("config", "./.go-pherson", "Path to version config file directory")
	flag.Parse()

	// Retrieve command
	command := flag.Arg(0)

	var configPath string = *dirPath

	fmt.Println(configPath)
	fmt.Println(command)
}
