package dot

// TODO: The operating system should help define the templates for things like
// `dot.bashrc`, etc.
type OperatingSystem int

const (
	Debian OperatingSystem = iota
	Alpine
	Ubuntu
	Fedora
)

func (self OperatingSystem) String() string {
	switch self {
	case Alpine:
		return "alpine"
	case Ubuntu:
		return "ubuntu"
	case Fedora:
		return "fedora"
	default: // Debian
		return "debian"
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
