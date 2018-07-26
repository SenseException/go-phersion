package main

import (
	"github.com/SenseException/go-phersion/config"
	"gopkg.in/urfave/cli.v1"
	"os"
	"fmt"
	"github.com/SenseException/go-phersion/versioning"
)

func main() {
	// path to the go-phersion config file
	var configPath string

	// Check if config path exists
	configExists := func(configPath string) error {
		if ! config.Exists(configPath) {
			return cli.NewExitError("No project config was initialized. Use: go-phersion init", 1)
		}

		return nil
	}

	// Create cli API of go-phersion
	app := cli.NewApp()
	app.Name = "go-phersion"
	app.Author = "Claudio Zizza"
	app.Usage = "Managing versions in builds of projects"
	app.EnableBashCompletion = true
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Println("Unknown command: ", command)
	}


	// Define arguments of go-phersion
	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "config, c",
			Value: "./.go-pherson",
			Usage: "Path to version config file directory",
			Destination: &configPath,
		},
	}

	// TODO Move console logic outside of main.go
	typeHandler := func(versionType string, processType func(version *versioning.Version, versionType string), message string) error {
		err := configExists(configPath)
		if nil != err {
			return err
		}

		if "" == versionType {
			return cli.NewExitError("Version type argument is missing", 2)
		}

		version, err := config.Read(configPath)
		if nil != err {
			return err
		}

		processType(&version, versionType)

		err = config.Write(version, configPath)
		if nil != err {
			return err
		}

		fmt.Println("Version type", versionType, message)

		return nil
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
		{
			Name: "bump",
			Usage: "Increments the version of the project",
			Action: func(c *cli.Context) error {
				err := configExists(configPath)
				if nil != err {
					return err
				}

				return nil
			},
		},
		{
			Name: "add-type",
			Usage: "Adds a new version type, that will contain the version in a config file usable for your project",
			ArgsUsage: "version-type",
			Action: func(c *cli.Context) error {
				return typeHandler(c.Args().Get(0), func(version *versioning.Version, versionType string) {
					version.AddType(versionType)
				}, "was added")
			},
		},
		{
			Name: "remove-type",
			Usage: "Removed an existing version type from your configuration",
			ArgsUsage: "version-type",
			Action: func(c *cli.Context) error {
				return typeHandler(c.Args().Get(0), func(version *versioning.Version, versionType string) {
					version.RemoveType(versionType)
				}, "was removed")
			},
		},
	}

	app.Run(os.Args)
}