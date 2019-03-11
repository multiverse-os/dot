package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	dot "github.com/multiverse-os/dot-config"
	log "github.com/multiverse-os/log"
	terminal "github.com/multiverse-os/os/terminal"
)

// TODO: Fix `multiverse-os/log` so that space is not required before the
// message

func main() {
	terminal.PrintBanner(terminal.White("[")+terminal.Blue("Multiverse OS")+terminal.White(": ")+terminal.White("dot.config")+terminal.White("]  ")+terminal.Light("basic system provisioning/config management"), terminal.White("0.1.0"))

	var profileArgument string
	if len(os.Args) > 1 {
		profileArgument = strings.ToLower(os.Args[1])
	} else {
		log.FatalError(errors.New(" Profile argument required (e.g. default, dev)"))
		fmt.Println(terminal.Strong(" Usage:\n") + terminal.White("   dotconfig default") + terminal.White("\n   dotconfig dev ./profiles/dev.golang.yaml"))
		os.Exit(0)
	}

	switch profileArgument {
	case dot.DefaultProfile.String():
		env := dot.DefaultConfig()
		fmt.Println("Loading default profile, an environment with a single profile, more profiles could be merged into it...")
		fmt.Println("Number of profiles: ", len(env.Profiles))
		fmt.Println("First profile type is: ", env.Profiles[0].Type)

	case dot.DevelopmentProfile.String(), "dev":
		if len(os.Args) > 2 && len(os.Args[1]) > 1 {
			envPath := os.Args[2]
			fmt.Println("Checking if [", envPath, "] exists...")
			if _, err := os.Stat(envPath); os.IsNotExist(err) {
				log.FatalError(errors.New(" Development provision configuration profile not found"))
				fmt.Println(terminal.Strong(" Usage:") + terminal.White("\n   dotconfig dev profiles/dev.golang.yaml"))
				os.Exit(0)
			} else {
				env, err := dot.LoadEnvironment(envPath)
				if err != nil {
					log.FatalError(errors.New(" Failed to parse provision configuration, verify profile YAML and try again"))
				} else {
					if len(env.Profiles) > 0 {
						fmt.Println("number of profiles in env [", len(env.Profiles), "] ")
						fmt.Println("Successfully loaded the configuration, now just neeed to iterate through each profile and provision based on it: ", env)
						fmt.Println("Loading default profile, an environment with a single profile, more profiles could be merged into it...")
						fmt.Println("Number of profiles: ", len(env.Profiles))
						fmt.Println("First profile type is: ", env.Profiles[0].Type)
					} else {
						log.FatalError(errors.New(" Failed to parse provision configuration, must have at least one profile"))
					}
				}
			}
		} else {
			log.FatalError(errors.New(" Must include developer provision configuration path"))
			fmt.Println(terminal.Strong(" Usage:") + terminal.White("\n   dotconfig dev profiles/dev.golang.yaml"))
		}
	default:
		log.FatalError(errors.New(" Invalid profile [available profiles: default, dev]"))
		fmt.Println(terminal.Strong(" Usage:\n") + terminal.White("   dotconfig default") + terminal.White("\n   dotconfig dev ./profiles/dev.golang.yaml"))
		os.Exit(0)
	}
}
