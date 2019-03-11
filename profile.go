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

type Environment struct {
	OS       string    `yaml:"os"`
	Version  string    `yaml:"version"`
	Profiles []Profile `yaml:"profiles"`
}

type Profile struct {
	Type                string                 `yaml:"type"`
	Subtype             string                 `yaml:"subtype"`
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
		OS:      "debian",
		Version: Version{Major: 9, Minor: 7, Patch: 0}.String(),
		Profiles: []Profile{
			Profile{
				Type: "default",
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

func LoadEnvironment(path string) (env *Environment, err error) {
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
