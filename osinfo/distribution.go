package osinfo

import (
	"path/filepath"
	"strings"
)

type Distribution int

func (self Distribution) OSFiles() []string {
	filenames, err := filepath.Glob(("osinfo/db/os/" + self.URL() + "/*"))
	if err != nil {
		panic(err)
	}
	return filenames
}

const (
	AlpineLinux Distribution = iota
	AltLinux
	Apple
	ArchLinux
	Android
	Asianux
	CentOS
	Debian
	DragonflyBSD
	EndlessOS
	FedoraProject
	FreeBSD
	FreeDOS
	Gentoo
	Gnome
	HaikuOS
	Mageia
	Mandriva
	Microsoft
	NetBSD
	Novell
	OpenBSD
	OpenSUSE
	Oracle
	PureOS
	Redhat
	ScientificLinux
	Sun
	SUSE
	System76
	Ubuntu
	VoidLinux
)

// Aliasing to make the API easier to use
const (
	Alpine = AlpineLinux
	Alt    = AltLinux
	Arch   = ArchLinux
	Fedora = FedoraProject
	Void   = VoidLinux
)

func (self Distribution) String() string {
	switch self {
	case AlpineLinux:
		return "alpine"
	case AltLinux:
		return "alt"
	case Apple:
		return "apple"
	case ArchLinux:
		return "arch"
	case Android:
		return "android"
	case Asianux:
		return "asianux"
	case CentOS:
		return "centos"
	case Debian:
		return "debian"
	case DragonflyBSD:
		return "dragonflybsd"
	case EndlessOS:
		return "endlessos"
	case FedoraProject:
		return "fedora"
	case FreeBSD:
		return "freebsd"
	case FreeDOS:
		return "freedos"
	case Gentoo:
		return "gentoo"
	case Gnome:
		return "gnome"
	case HaikuOS:
		return "haikuos"
	case Mageia:
		return "mageia"
	case Mandriva:
		return "mandriva"
	case Microsoft:
		return "miscrosoft"
	case NetBSD:
		return "netbsd"
	case Novell:
		return "novell"
	case OpenBSD:
		return "openbsd"
	case OpenSUSE:
		return "opensuse"
	case Oracle:
		return "oracle"
	case PureOS:
		return "pureos"
	case Redhat:
		return "redhat"
	case ScientificLinux:
		return "scientificlinux"
	case Sun:
		return "sun"
	case SUSE:
		return "suse"
	case System76:
		return "system76"
	case VoidLinux:
		return "void"
	default: // Ubuntu:
		return "ubuntu"
	}
}

func MarshalDistribution(name string) Distribution {
	switch strings.ToLower(name) {
	case AlpineLinux.String():
		return AlpineLinux
	case AltLinux.String():
		return AltLinux
	case Apple.String():
		return Apple
	case ArchLinux.String():
		return ArchLinux
	case Android.String():
		return Android
	case Asianux.String():
		return Asianux
	case CentOS.String():
		return CentOS
	case Debian.String():
		return Debian
	case DragonflyBSD.String():
		return DragonflyBSD
	case EndlessOS.String():
		return EndlessOS
	case FedoraProject.String():
		return FedoraProject
	case FreeBSD.String():
		return FreeBSD
	case FreeDOS.String():
		return FreeDOS
	case Gentoo.String():
		return Gentoo
	case Gnome.String():
		return Gnome
	case HaikuOS.String():
		return HaikuOS
	case Mageia.String():
		return Mageia
	case Mandriva.String():
		return Mandriva
	case Microsoft.String():
		return Microsoft
	case NetBSD.String():
		return NetBSD
	case Novell.String():
		return Novell
	case OpenBSD.String():
		return OpenBSD
	case OpenSUSE.String():
		return OpenSUSE
	case Oracle.String():
		return Oracle
	case PureOS.String():
		return PureOS
	case Redhat.String():
		return Redhat
	case ScientificLinux.String():
		return ScientificLinux
	case Sun.String():
		return Sun
	case SUSE.String():
		return SUSE
	case System76.String():
		return System76
	case VoidLinux.String():
		return VoidLinux
	default: // Ubuntu
		return Ubuntu
	}
}

func (self Distribution) URL() string {
	switch self {
	case AlpineLinux:
		return "alpinelinux.org"
	case AltLinux:
		return "altlinux.org"
	case Android:
		return "android-x86.org"
	case Apple:
		return "apple.com"
	case ArchLinux:
		return "archlinux.org"
	case Asianux:
		return "asianux.com"
	case CentOS:
		return "centos.org"
	case Debian:
		return "debian.org"
	case DragonflyBSD:
		return "dragonflybsd.org"
	case EndlessOS:
		return "endlessos.org"
	case FedoraProject:
		return "fedoraproject.org"
	case FreeBSD:
		return "freebsd.org"
	case FreeDOS:
		return "freedos.org"
	case Gentoo:
		return "gentoo.org"
	case Gnome:
		return "gnome.org"
	case HaikuOS:
		return "haiku-os.org"
	case Mageia:
		return "mageia.org"
	case Mandriva:
		return "mandriva.com"
	case Microsoft:
		return "miscrosoft.com"
	case NetBSD:
		return "netbsd.org"
	case Novell:
		return "novell.com"
	case OpenBSD:
		return "openbsd.org"
	case OpenSUSE:
		return "opensuse.org"
	case Oracle:
		return "oracle.com"
	case PureOS:
		return "pureos.net"
	case Redhat:
		return "redhat.com"
	case ScientificLinux:
		return "scientificlinux.org"
	case Sun:
		return "sun.com"
	case SUSE:
		return "suse.com"
	case System76:
		return "system76.com"
	case VoidLinux:
		return "voidlinux.org"
	default: // Ubuntu:
		return "ubuntu.com"
	}
}
