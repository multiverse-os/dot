package dot

var PackageNames []string
var Packages []Package

type Package struct {
	name string
	installed bool
}

func (self *Package) Install() error {
	fmt.Println(`  Installing package: '` + self.name + `'`)
	return bash(`sudo apt-get install -y ` + self.name)
}

func (self *Package) Uninstall() error {
	fmt.Println("  Removing package: '" + self.name + "'")
	return bash("sudo apt-get -y remove" + self.name)
}
