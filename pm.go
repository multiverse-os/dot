package dot

import (
	"strings"
)

//type PackageManager interface {
//	Install() (bool, error)
//	Uninstall() (bool, error)
//	Package() string
//	Dependencies() []string
//	InstallDependencies() (bool, error)
//	Configs() ([]string, error)
//	InstallConfigs() (bool, error)
//	PostInstallCommands() []string
//	RunPostInstallCommands() (bool, error)
//	Installed() bool
//}

type PackageManager int

const (
	Apt PackageManager = iota
	Apk
	Dnf
)

func (self PackageManager) String() string {
	switch self {
	case Apk:
		return "apk"
	case Dnf:
		return "dnf"
	default: // Apt
		return "apt"
	}
}

func (self PackageManager) Install() string {
	switch self {
	case Apk:
		return "apk add"
	case Dnf:
		return "dnf install"
	default: // Apt
		return "DEBIAN_FRONTEND=noninteractive apt-get install -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y --no-install-recommends1"
	}
}

func (self PackageManager) InstallPackage(pkg string) error {
	return terminal(self.Install() + ` ` + pkg)
}

func (self PackageManager) SudoInstallPackage(pkg string) error {
	return terminal(`sudo ` + self.Install() + ` ` + pkg)
}

func (self PackageManager) InstallPackages(pkgs []string) error {
	return terminal(self.Install() + ` ` + strings.Join(pkgs, " "))
}

func (self PackageManager) SudoInstallPackages(pkgs []string) error {
	return terminal(`sudo ` + self.Install() + ` ` + strings.Join(pkgs, " "))
}

func (self PackageManager) Remove() string {
	switch self {
	case Apk:
		return "apk rm"
	case Dnf:
		return "dnf remove"
	default: // Apt
		return "DEBIAN_FRONTEND=noninteractive apt-get remove -y"
	}
}

func (self PackageManager) RemovePackage(pkg string) error {
	return terminal(self.Remove() + ` ` + pkg)
}

func (self PackageManager) SudoRemovePackage(pkg string) error {
	return terminal(`sudo ` + self.Remove() + ` ` + pkg)
}

func (self PackageManager) RemovePackages(pkgs []string) error {
	return terminal(self.Remove() + ` ` + strings.Join(pkgs, " "))
}

func (self PackageManager) SudoRemovePackages(pkgs []string) error {
	return terminal(`sudo ` + self.Remove() + ` ` + strings.Join(pkgs, " "))
}

func (self PackageManager) Installed() string {
	switch self {
	case Apk:
		return "apk list"
	case Dnf:
		return "dnf list installed"
	default: // Apt
		return "apt list --installed"
	}
}

func (self PackageManager) AddRepository() {
	// TODO: Add lines to repository
	// TODO: Edit repository line
	// TODO: Reset repository to default
}

func (self PackageManager) Update() string {
	switch self {
	case Apk:
		return "apk update"
	case Dnf:
		return "dnf update"
	default: // Apt
		return "DEBIAN_FRONTEND=noninteractive apt-get update -y"
	}
}

func (self PackageManager) Upgrade() string {
	switch self {
	case Apk:
		return "apk update"
	case Dnf:
		return "dnf update"
	default: // Apt
		return "DEBIAN_FRONTEND=noninteractive apt-get upgrade -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y"
	}
}

func (self PackageManager) DistUpgrade() string {
	switch self {
	case Dnf:
		return "dnf distro-sync"
	default: // Apt
		return "DEBIAN_FRONTEND=noninteractive apt-get dist-upgrade -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y"
	}
}
