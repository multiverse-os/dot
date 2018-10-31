package dot

import config "github.com/multiverse-os/dot-manager/config"

type UserProfile int

const (
	DEFAULT_PROFILE UserProfile = iota
	DEVELOPER
)

var pkgNamesToAdd PackageNames = PackageNames{"curl", "neovim", "vim", "vim-gocomplete", "golang", "make", "cmake", "build-essential"}

var pkgNamesToDel PackageNames = PackageNames{"nano"}

func (self UserProfile) String() string {
	switch self {
	case DEVELOPER:
		return "Developer"
	default: // DEFAULT_PROFILE
		return "Default"
	}
}

func initPackages(pkgNames PackageNames) Packages {
	var pkgs Packages
	for _, pkgName := range pkgNames {
		pkgs = append(pkgs, Package{name: pkgName})
		// TODO would be better to check if pkg is installed here
		// something like pkg.VerifyInstalled()
		// should be distro agnostic
	}
	return pkgs
}

func InstallProfile(userProfile UserProfile) []error {
	var installErrs []error

	// TODO: Divide pkgs by profile
	pkgsToAdd := initPackages(pkgNamesToAdd)
	pkgsToDel := initPackages(pkgNamesToAdd)

	for _, pkg := range pkgsToAdd {
		// TODO: check p.installed
		err := pkg.Install()
		if err != nil {
			installErrs = append(installErrs, err)
		} else {
			pkg.installed = true
		}
	}

	for _, pkg := range pkgsToDel {
		err := pkg.Uninstall()
		if err != nil {
			installErrs = append(installErrs, err)
		} else {
			pkg.installed = false
		}
	}

	switch userProfile {
	case DEVELOPER:
		config.Install("neovim")
	default: // Default Profile
	}
	config.Install("terminal.Bashrc")
	config.Install("git")

	successMsg := "Sucessfully installed or removed listed packages"
	return installErrs
}
