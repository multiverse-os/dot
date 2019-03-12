package dot

// TODO: The operating system should help define the templates for things like
// `dot.bashrc`, etc.
type OperatingSystem int

const (
	Alpine OperatingSystem = iota
	Debian
	Fedora
	Ubuntu
)

func (self OperatingSystem) String() string {
	switch self {
	case Alpine:
		return "alpine"
	case Fedora:
		return "fedora"
	case Ubuntu:
		return "ubuntu"
	default: // Debian
		return "debian"
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
	default:
		return Debian
	}
}

func (self OperatingSystem) PackageManager() PackageManager {
	switch self {
	case Alpine:
		return Apk
	case Fedora:
		return Dnf
	default: // Debian, Ubuntu
		return Apt
	}
}
