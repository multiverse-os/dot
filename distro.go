package dot

var OSInfoDBRepository = "https://gitlab.com/libosinfo/osinfo-db"

type Distribution struct {
	Name    string  `yaml:"Name"`
	Version Version `yaml:"version"`
}
