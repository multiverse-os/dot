package dot

import "fmt"

//"github.com/go-cmd/cmd"

//e := executor.New(exec.Command("/bin/sh", "echo hello"))
//e.Start() // start

func (self Environment) Provision() (errors []error) {
	os := MarshallOS(self.OS)
	for _, profile := range self.Profiles {
		errors = append(errors, profile.Provision(os.PackageManager()))
	}
	return errors
}

func (self Profile) Provision(pm PackageManager) error {
	terminal(`echo "test"`)

	var pkgs []string
	for _, pkg := range self.Packages.Install {
		pkgs = append(pkgs, pkg)
	}
	fmt.Println("Packages to install: ", pkgs)

	return nil
}
