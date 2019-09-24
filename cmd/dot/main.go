package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
	dot "github.com/multiverse-os/dot"
	osinfo "github.com/multiverse-os/dot/osinfo"
)

func main() {
	// TODO: Ideally will just merge in or basically forward all existing `ip
	// commands` and use the command name `ip`, so basically it will function as a
	// way of adding in more consistent functionality (like adding JSON output to
	// every command, and providing more functionality, while keeping all the
	// original functionality and expected usage)

	cmd := cli.New(&cli.CLI{
		Name:    "dot",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 0},
		//Usage:   "Specify a command, and one ip address or more",
		Commands: []cli.Command{
			cli.Command{
				Name:    "provision",
				Aliases: []string{"p"},
				Usage:   "provision using specified configuration",
				Action: func(c *cli.Context) error {
					environment := dot.DefaultEnvironment()
					fmt.Println("config:", environment)

					info := osinfo.Load(osinfo.Debian)
					fmt.Println("osInfo:", info)

					//config.Provision()

					return nil
				},
			},
		},
		//DefaultAction: func(context *cli.Context) error {
		//	cli.RenderHelp()
		//	return nil
		//},
	})

	cmd.Run(os.Args)
}
