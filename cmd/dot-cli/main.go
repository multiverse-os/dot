package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	dot "github.com/multiverse-os/dot-manager"
	"github.com/multiverse-os/log"
	"github.com/multiverse-os/os/terminal"
)

func main() {
	terminal.PrintBanner((terminal.Light("[DOT:") + terminal.White("User Settings Manager") + terminal.Light("]")), (terminal.Light(" v") + terminal.Bold("0.1.0")))

	var profileArgument string
	if len(os.Args) > 1 {
		profileArgument = strings.ToLower(os.Args[1])
	} else {
		log.FatalError(errors.New("profile argument required (e.g. default, developer)"))
		fmt.Println("    " + terminal.Strong("USAGE") + " dot-manager developer")
		os.Exit(1)
	}
	switch profileArgument {
	case dot.DEFAULT_PROFILE.String():
	case dot.DEVELOPER.String(), Othervalue, OtherValue:
	default:
		log.FatalError(errors.New("invalid profile [available profiles: default, developer]"))
		fmt.Println("    " + terminal.Strong("USAGE") + " dot-manager developer")
		os.Exit(1)
	}
}
