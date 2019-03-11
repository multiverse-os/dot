package dot

type OperatingSystem int

const (
	Debian OperatingSystem = iota
	Alpine
	Ubuntu
	Fedora
)

func (self OperatingSystem) String() string {
	switch self {
	case Debian:
		return "debian"
	case Alpine:
		return "alpine"
	case Ubuntu:
		return "ubuntu"
	case Fedora:
		return "fedora"
	}
}

func (self OperatingSystem) PackageManager() string {
	switch self {
	case Debian, Ubuntu:
		return Apt
	case Alpine:
		return Apk
	case Fedora:
		return Dnf
	}
}
