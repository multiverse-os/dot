package dot

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
		return "DEBIAN_FRONTEND=noninteractive apt-get install -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y"
	}
}

func (self PackageManager) InstallPackage(pkg string) {
	terminal(self.Install() + ` ` + pkg)
}

func (self PackageManager) InstallPackages(pkgs []string) {
	terminal(self.Install() + ` ` + pkg)
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
