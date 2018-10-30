package dot

type userProfile int

const (
	defaultProfile userProfile = iota
	developer
)

func InstallProfile(userProfile UserProfile) (string, error) {
	// TODO: Divide packages by profile
	InstallPackages(PackageNames{"curl", "neovim", "vim", "vim-gocomplete", "golang", "make", "cmake", "build-essential"})
	RemovePackages(PackageNames{"nano"})

	case developerProfile:
		InstallConfig("neovim")
	default: // Default Profile
	}
	InstallConfig("bashrc")
	InstallConfig("git")
}

func InstallPackage(packages string) (string, error) {
	fmt.Println(`  Installing package: '` + fmt.Sprintf(packages) + `'`)
	return bash(`sudo apt-get install -y ` + fmt.Sprintf(packages))
}

func InstallPackages(packages PackageNames) (string, error) {
	fmt.Println(`  Installing package: '` + fmt.Sprintf(packages) + `'`)
	return bash(`sudo apt-get install -y ` + fmt.Sprintf(packages))
}

