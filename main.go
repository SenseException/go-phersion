package main

import (
	"github.com/SenseException/go-phersion/config"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	// path to the go-phersion config file
	var configPath string

	// Create cli API of go-phersion
	app := cli.NewApp()
	app.Name = "go-phersion"
	app.Author = "Claudio Zizza"
	app.Description = "Managing versions in builds of projects"


	// Define arguments of go-phersion
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config, c",
			Value: "./.go-pherson",
			Usage: "Path to version config file directory",
			Destination: &configPath,
		},
	}

	// Create commands of go-phersion
	app.Commands = []cli.Command {
		{
			Name: "init",
			Usage: "Initialize version for a new project",
			Action: func(c *cli.Context) error {
				if ! config.Exists(configPath) {
					config.Init(configPath)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}