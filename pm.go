package dot

type PackageManager int

const (
	Apt PackageManager = iota
	Apk
	Dnf
)

func (self PackageManager) String() string {
	switch self {
	case Apt:
		return "apt"
	case Apk:
		return "apk"
	case Dnf:
		return "dnf"
	}
}

func (self PackageManager) Install() string {
	switch self {
	case Apt:
		return "DEBIAN_FRONTEND=noninteractive apt-get install -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y"
	case Apk:
		return "apk add"
	case Dnf:
		return "dnf install"
	}
}

func (self PackageManager) Remove() string {
	switch self {
	case Apt:
		return "DEBIAN_FRONTEND=noninteractive apt-get remove -y"
	case Apk:
		return "apk rm"
	case Dnf:
		return "dnf remove"
	}
}

func (self PackageManager) Update() string {
	switch self {
	case Apt:
		return "DEBIAN_FRONTEND=noninteractive apt-get update -y"
	case Apk:
		return "apk update"
	case Dnf:
		return "dnf update"
	}
}

func (self PackageManager) Upgrade() string {
	switch self {
	case Apt:
		return "DEBIAN_FRONTEND=noninteractive apt-get upgrade -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y"
	case Apk:
		return "apk update"
	case Dnf:
		return "dnf update"
	}
}

func (self PackageManager) DistUpgrade() string {
	switch self {
	case Apt:
		return "DEBIAN_FRONTEND=noninteractive apt-get dist-upgrade -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y"
	case Dnf:
		return "dnf distro-sync"
	}
}

func (self PackageManager) Installed() string {
	switch self {
	case Apt:
		return "apt list --installed"
	case Apk:
		return "apk list"
	case Dnf:
		return "dnf list installed"
	}
}
