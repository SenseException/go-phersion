package main

import (
	"github.com/SenseException/go-phersion/config"
	"gopkg.in/urfave/cli.v1"
	"os"
	"fmt"
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


	// Define arguments of go-phersion
	app.Flags = []cli.Flag {
		cli.StringFlag {
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
				err := configExists(configPath)
				if nil != err {
					return err
				}

				var versionType string = c.Args().Get(0)
				if "" == versionType {
					return cli.NewExitError("Version type argument is missing", 2)
				}

				version, err := config.Read(configPath)
				if nil != err {
					return err
				}

				version.AddType(versionType)

				err = config.Write(version, configPath)
				if nil != err {
					return err
				}

				fmt.Println("Version type", versionType, "was added")

				return nil
			},
		},
		{
			Name: "remove-type",
			Usage: "Removed an existing version type from your configuration",
			ArgsUsage: "version-type",
			Action: func(c *cli.Context) error {
				err := configExists(configPath)
				if nil != err {
					return err
				}

				var versionType string = c.Args().Get(0)
				if "" == versionType {
					return cli.NewExitError("Version type argument is missing", 2)
				}

				version, err := config.Read(configPath)
				if nil != err {
					return err
				}

				version.RemoveType(versionType)

				err = config.Write(version, configPath)
				if nil != err {
					return err
				}

				fmt.Println("Version type", versionType, "was removed")

				return nil
			},
		},
	}

	app.Run(os.Args)
}