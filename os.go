package dot

// TODO: The operating system should help define the templates for things like
// `dot.bashrc`, etc.
type OperatingSystem int

type Distribution struct {
	OS             OperatingSystem
	Version        Version `yaml:"version"`
	PackageManager PackageManager
}

func LoadDistribution(name string) *Distribution {
	return &Distribution{
		OS:      Debian,
		Version: Version{Major: 9, Minor: 9, Patch: 0},
	}
}

const (
	Alpine OperatingSystem = iota
	Android
	Arch
	CentOS
	Debian
	Fedora
	Gentoo
	OpenSUSE
	Redhat
	Ubuntu
)

func (self OperatingSystem) String() string {
	switch self {
	case Alpine:
		return "alpine"
	case Android:
		return "android"
	case Arch:
		return "arch"
	case CentOS:
		return "centos"
	case Debian:
		return "debian"
	case Fedora:
		return "fedora"
	case Gentoo:
		return "gentoo"
	case OpenSUSE:
		return "opensuse"
	case Redhat:
		return "redhat"
	default: // Ubuntu
		return "ubuntu"
	}
}

func MarshalOS(os string) OperatingSystem {
	switch os {
	case Alpine.String():
		return Alpine
	case Fedora.String():
		return Fedora
	case Ubuntu.String():
		return Ubuntu
	default: // Debian.String()
		return Debian
	}
}
