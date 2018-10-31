package dot

import (
	"fmt"

	"github.com/multiverse-os/os/terminal"
)

type Packages []Package
type PackageNames []string

type Package struct {
	name      string
	installed bool
}

func (self Package) Install() error {
	fmt.Println(`  Installing package: '` + self.name + `'`)
	_, err := terminal.Bash(`sudo apt install -y ` + self.name)
	return err
}

func (self Package) Uninstall() error {
	fmt.Println("  Removing package: '" + self.name + "'")
	_, err := terminal.Bash("sudo apt -y remove" + self.name)
	return err
}

func (self Package) VerifyInstalled() (bool, error) {
	// TODO: distro agnostic method of checking if a package is installed
	return false, nil
}
