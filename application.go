package dot

// TODO: YAML config for each then the profile can call in thse YAML configs for
// the application installation. Like neovim complete installation can be
// defined and then that can be called into a profile for setting up development
// environment.

// Then ideally mix and match profiles to setup a custom environment
type ApplicationManager interface {
	Install() (bool, error)
	Uninstall() (bool, error)
	Package() string
	Dependencies() []string
	InstallDependencies() (bool, error)
	ConfigFiles() ([]string, error)
	InstallConfigFiles() (bool, error)
	PostInstallCommands() []string
	RunPostInstallCommands() (bool, error)
	Installed() bool
}

// TODO: Make package an object, then we can just ask if its installed, etc.

// TODO: Additionally commands should be object
type Application struct {
	Name                string
	Package             string
	DependentPackages   []string
	ConfigFiles         []ConfigFile
	PostInstallCommands []string
}
