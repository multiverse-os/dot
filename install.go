package dot

import (
	"fmt"

	"github.com/multiverse-os/os/terminal"
)

type UserProfile int

const (
	DEFAULT_PROFILE UserProfile = iota
	DEVELOPER
)

func (self UserProfile) String() string {
	switch self {
	case DEVELOPER:
		return "Developer"
	default: // DEFAULT_PROFILE
		return "Default"
	}
}

func InstallProfile(userProfile UserProfile) (string, error) {
	// TODO: Divide packages by profile
	InstallPackages(PackageNames{"curl", "neovim", "vim", "vim-gocomplete", "golang", "make", "cmake", "build-essential"})
	RemovePackages(PackageNames{"nano"})
	switch userProfile {
	case DEVELOPER:
		InstallConfig("neovim")
	default: // Default Profile
	}
	InstallConfig("bashrc")
	InstallConfig("git")
}

func InstallPackage(packages string) (string, error) {
	fmt.Println(`  Installing package: '` + fmt.Sprintf(packages) + `'`)
	return terminal.Bash(`sudo apt-get install -y ` + fmt.Sprintf(packages))
}

func InstallPackages(packages PackageNames) (string, error) {
	fmt.Println(`  Installing package: '` + fmt.Sprintf(packages) + `'`)
	return bash(`sudo apt-get install -y ` + fmt.Sprintf(packages))
}
