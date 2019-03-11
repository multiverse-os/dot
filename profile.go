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
	DeveloperProfile
)

type DeveloperSubtype int

const (
	CLanguage DeveloperSubtype = iota // Default
	GoLanguage
	RustLanguage
	RubyLanguage
	PythonLanguage
)

type Profile struct {
	Type     ProfileType
	Language DeveloperSubtype
	Filename string
}

type EnvironmentFile struct {
	Profiles []ProfileConfiguration `yaml:"env"`
}

type ProfileConfig struct {
	Type            string `yaml:"type"`
	Subtype         string `yaml:"subtype"`
	OperatingSystem string `yaml:"os"`
	ProfilePackages struct {
		Install []string `yaml:"install"`
		Remove  []string `yaml:"remove"`
	} `yaml:"packages"`
	ProfileConfigFiles  []ProfileConfigInstall `yaml:"configurations"`
	PostInstallCommands []string               `yaml:"commands"`
}

type ProfileConfigInstall struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

func DefaultConfig() Config {
	return EnvironmentFile{
		Profiles: []ProfileConfig{
			ProfileConfiguration{
				Type:            "default",
				OperatingSystem: "debian",
				ProfilePackages: {
					Install: []string{"curl", "git", "vim"},
					Remove:  []string{"nano"},
				},
				ProfileConfigurationFiles: {
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

func LoadConfig(path string) (config *Config, err error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (self *Config) Save(path string) error {
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
