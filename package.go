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

func (self *Package) Install() error {
	fmt.Println(`  Installing package: '` + self.name + `'`)
	return terminal.Bash(`sudo apt-get install -y ` + self.name)
}

func (self *Package) Uninstall() error {
	fmt.Println("  Removing package: '" + self.name + "'")
	return terminal.Bash("sudo apt-get -y remove" + self.name)
}
