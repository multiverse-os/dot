package dot

import (
	"encoding/xml"
	"fmt"
	"os"
	//"io/ioutil"
)

var OSInfoDBRepository = "https://gitlab.com/libosinfo/osinfo-db"

// TODO: Build a script to determine newest version and recreate this map
var xmlFile = map[string]string{
	"Debian":   "./osinfo/os/debian.org/debian-10.xml.in",
	"Ubuntu":   "./osinfo/os/ubuntu.com/ubuntu-19.04.xml.in",
	"Redhat":   "./osinfo/os/redhat.com/rhl-9.xml.in",
	"CentOS":   "./osinfo/os/centos.org/centos-7.0.xml.in",
	"Fedora":   "./osinfo/os/fedoraproject.org/fedora-30.xml.in",
	"Gentoo":   "./osinfo/os/gentoo.org/gentoo-rolling.xml.in",
	"Alpine":   "./osinfo/os/alpinelinux.org/alpinelinux-3.8.xml.in",
	"OpenSUSE": "./osinfo/os/opensuse.org/opensuse-42.3.xml.in",
	"Android":  "./osinfo/os/android-x86.org/android-x86-8.1.xml.in",
	"Arch":     "./osinfo/os/archlinux.org/archlinux-rolling.xml.in",
}

func LoadOSInfo(name string) (osInfo *OSInfo) {

	root := &Node{}
	err := NewDecoder(r, ps...).Decode(root)
	if err != nil {
		return nil, err
	}

	//xmlBytes, err := ioutil.ReadFile(distributionInfo[distributionName])
	osXML, err := os.Open(xmlFile[name])
	decoder := xml.NewDecoder(osXML)
	//fmt.Println("len of xmlBytes:", len(xmlBytes))
	err = decoder.Decode(&osInfo)
	//err = xml.Unmarshal(xmlBytes, &os)
	if err != nil {
		panic(err)
	}
	fmt.Println("osInfo:", osInfo)
	return osInfo
}

type OSInfo struct {
	OperatingSystem OS `xml:"libosinfo"`
}

type OS struct {
	Id string `xml:"id,attr"`
	//ReleaseShortID string `xml:"os>short-id"`
	Version string `xml:"version"`
	//Vendor         string      `xml:"_vendor"`
	//Family         string      `xml:"family"`
	//Name           string      `xml:"distro"`
	//ReleaseName    string      `xml:"codename"`
	//UpgradesURL    string      `xml:"upgrade"`
	//ReleaseDate    string      `xml:"release-date"`
	//Devices        []Device    `xml:"devices,device"`
	//Resources      []Resources `xml:"resources"`
	//Variants       []Variant   `xml:"variant"`
	//Media          []Media     `xml:"media"`
	//ArchTrees      []ArchTree  `xml:"tree"`
	//Images         []Image     `xml:"image"`
}

type Device struct {
	URL string `xml:"attr,id"`
}

type Resources struct {
	Arch        string `xml:"arch,attr"`
	Minimum     Specs  `xml:"minimum"`
	Recommended Specs  `xml:"recommended"`
}

type Specs struct {
	CPU     int `xml:"cpu"`
	RAM     int `xml:"ram"`
	Storage int `xml:"storage"`
}

type Variant struct {
	ID   string `xml:"variant,attr"`
	Name string `xml:"_name"`
}

type Media struct {
	Arch      string `xml:"media,attr"`
	MediaType string `xml:"variant,attr"`
	URL       string `xml:"url"`
	ISO       string `xml:"iso,volume-id"`
	Kernel    string `xml:"kernel"`
	Initrd    string `xml:"initrd"`
}

type ArchTree struct {
	Arch   string `xml:"tree,attr"`
	URL    string `xml:"url"`
	Kernel string `xml:"kernel"`
	Initrd string `xml:"initrd"`
}

type Image struct {
	Arch   string `xml:"arch,attr"`
	Format string `xml:"format,attr"`
	Cloud  bool   `xml:"cloud-init,attr"`
}

type Installer struct {
	Preseeds []Script `xml:"script"`
}

type Script struct {
	PreseedURL string `xml:"id,attr"`
}
