package main

import (
	"fmt"
	"os"

	"libs/cli"
)

func main() {
	// TODO: Look for the default gravity config, if it does not exist write it
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "Gravity Build System"
	app.Description = "A Go language build system inspired by the Bundler dependency management system for ruby."

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initializes a Gravity project in the current directory",
			Action: func(c *cli.Context) error {
				fmt.Println("Initializing...")
				// TODO: Write PkgFile if it does not already exist
				return nil
			},
		},
	}

	app.Run(os.Args)
}
