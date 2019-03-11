package dot

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ProfileType int

const (
	DefaultProfile ProfileType = iota
	DevelopmentProfile
)

func (self ProfileType) String() string {
	switch self {
	case DevelopmentProfile:
		return "developer"
	default:
		return "default"
	}
}

type Language int

const (
	C Language = iota // Default
	Go
	Rust
	Ruby
	Python
)

func (self Language) String() string {
	switch self {
	case Go:
		return "Go"
	case Rust:
		return "Rust"
	case Ruby:
		return "Ruby"
	case Python:
		return "Python"
	default:
		return "C"
	}
}

type Environment struct {
	Profiles []Profile `yaml:"env"`
}

type Profile struct {
	Type                string                 `yaml:"type"`
	Subtype             string                 `yaml:"subtype"`
	OS                  string                 `yaml:"os"`
	Packages            ProfilePackages        `yaml:"packages"`
	ConfigFiles         []ProfileConfigInstall `yaml:"configurations"`
	PostInstallCommands []string               `yaml:"commands"`
}

type ProfilePackages struct {
	Install []string `yaml:"install"`
	Remove  []string `yaml:"remove"`
}

type ProfileConfigInstall struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

func DefaultConfig() Environment {
	return Environment{
		Profiles: []Profile{
			Profile{
				Type: "default",
				OS:   "debian",
				Packages: ProfilePackages{
					Install: []string{"curl", "git", "vim"},
					Remove:  []string{"nano"},
				},
				ConfigFiles: []ProfileConfigInstall{
					ProfileConfigInstall{
						From: "dot.bashrc",
						To:   "~/.bashrc",
					},
					ProfileConfigInstall{
						From: "dot.gitconfig",
						To:   "~/.gitconfig",
					},
				},
				PostInstallCommands: []string{"echo \"dot.config default profile install complete\""},
			},
		},
	}
}

func LoadConfig(path string) (env *Environment, err error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &env)
	if err != nil {
		return nil, err
	}
	return env, nil
}

func (self *Environment) Save(path string) error {
	configPath, _ := filepath.Split(path)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return err
	} else {
		yamlData, err := yaml.Marshal(&self)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(path, yamlData, 0600)
	}
}
