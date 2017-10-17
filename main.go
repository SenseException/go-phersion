package main

import "fmt"
import "os"
import "flag"

func main() {
	// Define arguments
	fileFlag := flag.String("file", "./.go-pherson", "Path to version config file directory")
	flag.Parse()

	// Retrieve command
	command := flag.Arg(0)

	var filePath string = *fileFlag

	os.MkdirAll(filePath, 0644)

	fmt.Println(filePath)
	fmt.Println(command)
}
