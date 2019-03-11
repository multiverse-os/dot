package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	dot "github.com/multiverse-os/dot-manager"
	log "github.com/multiverse-os/log"
	terminal "github.com/multiverse-os/os/terminal"
)

func main() {
	terminal.PrintBanner((terminal.Light("[dot.config:") + terminal.White("basic provisioning system") + terminal.Light("]")), (terminal.Light(" v") + terminal.Bold("0.1.0")))

	var profileArgument string
	if len(os.Args) > 1 {
		profileArgument = strings.ToLower(os.Args[1])
	} else {
		log.FatalError(errors.New("Profile argument required (e.g. default, developer)"))
		fmt.Println("    " + terminal.Strong("USAGE") + " dot.config developer")
		os.Exit(1)
	}

	switch profileArgument {
	case dot.DefaultProfile.String():
	case dot.DeveloperProfile.String(), Othervalue, OtherValue:
	default:
		log.FatalError(errors.New("Invalid profile [available profiles: default, developer]"))
		fmt.Println("    " + terminal.Strong("USAGE") + " dot.config developer")
		os.Exit(1)
	}
}
