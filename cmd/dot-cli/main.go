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

const (
	COMMAND_NOT_FOUND = "Must specify one of the available actions: [\"install\"]"
	CONFIG_NOT_FOUND  = "Failed to incldude provisioning configuration YAML file containing profiles"
	INVALID_CONFIG    = "Provision configuration parse failed, verify YAML is valid and try again"
	PROFILE_NOT_FOUND = "Provision configuration parse failed, must have at least one profile"
)

func Usage() {
	fmt.Println(terminal.Strong(" Usage:\n") + terminal.White("   dotconfig install default") + terminal.White("\n   dotconfig install ./profiles/dev.golang.yaml"))
}

func Title() string {
	return (terminal.White("[") + terminal.Blue("Multiverse OS") + terminal.White(": ") + terminal.White("dot.config") + terminal.White("]  ") + terminal.Light("basic system provisioning/config management"))
}

func main() {
	terminal.PrintBanner(Title(), terminal.White("0.1.0"))
	var profileArgument string
	if len(os.Args) > 1 {
		profileArgument = strings.ToLower(os.Args[1])
	}
	switch profileArgument {
	case "install", "i":
		if len(os.Args) > 2 && len(os.Args[1]) > 1 {
			envPath := os.Args[2]
			fmt.Println("Checking if [", envPath, "] exists...")
			if _, err := os.Stat(envPath); os.IsNotExist(err) {
				log.FatalError(errors.New(" " + CONFIG_NOT_FOUND))
			} else {
				env, err := dot.LoadEnvironment(envPath)
				if err != nil {
					log.FatalError(errors.New(" " + INVALID_CONFIG))
				} else {
					if len(env.Profiles) > 0 {
						fmt.Println("number of profiles in env [", len(env.Profiles), "] ")
						fmt.Println("First profile type is: ", env.Profiles[0].Type)
						errors := env.Provision()
						fmt.Println("Encountered [", len(errors), "] while attempting to provision with the configuration.")
					} else {
						log.FatalError(errors.New(" " + PROFILE_NOT_FOUND))
					}
				}
			}
		} else {
			log.FatalError(errors.New(" " + COMMAND_NOT_FOUND))
		}
	default:
		log.FatalError(errors.New(" " + COMMAND_NOT_FOUND))
	}
	Usage()
}
