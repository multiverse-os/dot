package osinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var OSInfoDBRepository = "https://gitlab.com/libosinfo/osinfo-db"

func Load(distro Distribution) (infoFiles []*InfoFile) {
	filenames := distro.OSFiles()

	for _, filename := range filenames {
		fmt.Println("filename:", filename)

		jsonBytes, _ := ioutil.ReadFile(filename)
		infoFile := &InfoFile{}
		err := json.Unmarshal(jsonBytes, &infoFile)
		if err != nil {
			fmt.Println("[error] error unmarshalling json data:", err)
		}

		infoFiles = append(infoFiles, infoFile)
	}

	fmt.Println("len(infoFiles):", len(infoFiles))
	fmt.Println("osInfo:", infoFiles)

	for _, file := range infoFiles {
		fmt.Println("file:", file)
		fmt.Println("file.LibOSInfo.OS.ReleaseVersion:")
		fmt.Println("rv:", file.LibOSInfo.OS.ReleaseVersion)
	}

	return infoFiles
}

type InfoFile struct {
	LibOSInfo struct {
		InfoVersion string `json:"-version"`
		OS          struct {
			ShortIDs        []string `json:"short-id"`
			ReleaseDate     string   `json:"release-date"`
			ReleaseName     string   `json:"codename"`
			ReleaseVersion  string   `json:"version"`
			Vendor          string   `json:"_vendor"`
			Family          string   `json:"linux"`
			URL             string   `json:"-id"`
			Name            string   `json:"name"`
			Distro          string   `json:"distro"`
			PreviousVersion struct {
				URL string `json:"-id"`
			} `json:"derives-from"`
			Variant []struct {
				Id   string `json:"-id"`
				Name string `json:"_name"`
			} `json:"variant"`
			Media []struct {
				Arch    string `json:"-arch"`
				Variant struct {
					Id string `json:"-id"`
				} `json:"variant"`
				URL string `json:"url"`
				ISO struct {
					VolumeID string `json:"volume-id"`
				} `json:"iso"`
				Kernel string `json:"kernel"`
				Initrd string `json:"initrd"`
			} `json:"media"`
			Tree []struct {
				Kernel string `json:"kernel"`
				Initrd string `json:"initrd"`
				Arch   string `json:"-arch"`
				URL    string `json:"url"`
			} `json:"tree"`
			Image []struct {
				Arch   string `json:"-arch"`
				Format string `json:"-format"`
				Cloud  string `json:"-cloud-init"`
				URL    string `json:"url"`
			} `json:"image"`
			Installer struct {
				Script []struct {
					Id string `json:"-id"`
				} `json:"script"`
			} `json:"installer"`
			Upgrades struct {
				Id string `json:"-id"`
			} `json:"upgrades"`
			Resources struct {
				Arch        string `json:"-arch"`
				Recommended struct {
					Storage string `json:"storage"`
					CPU     string `json:"cpu"`
					RAM     string `json:"ram"`
				} `json:"recommended"`
				Minimum struct {
					Storage string `json:"storage"`
					CPU     string `json:"cpu"`
					RAM     string `json:"ram"`
					CPUs    string `json:"n-cpus"`
				} `json:"minimum"`
			} `json:"resources"`
		} `json:"os"`
	} `json:"libosinfo"`
}

type ByVersion []*InfoFile

func (self ByVersion) Len() int      { return len(self) }
func (self ByVersion) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func (self ByVersion) Less(i, j int) bool {
	return self[i].LibOSInfo.OS.ReleaseVersion < self[j].LibOSInfo.OS.ReleaseVersion
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
