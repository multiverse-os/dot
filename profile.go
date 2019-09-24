package dot

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type ProfileType int

const (
	Default ProfileType = iota
	Multiverse
	Development
)

func (self ProfileType) String() string {
	switch self {
	case Development:
		return "developer"
	case Multiverse:
		return "multiverse"
	default: // Default
		return "default"
	}
}

// NOTE: Profile is a collection of configurations. These can be mixed and
// matched using Environment. They can be deifned in a git repository, or on the
// local disk. They have a category, type and optional subtype. This enables
// 		Category: Development
//    Type:     (Ruby|Golang|Rust|Python)
//    Subtype:  (Production|Development|...)
// The profile should include everything needed to reverse the installation
// process.
// If the uninstall require similar to post install, post uninstall commands
type Profile struct {
	Category            string          `yaml:"category"`
	Type                string          `yaml:"type"`
	Subtype             string          `yaml:"subtype"`
	Packages            ProfilePackages `yaml:"packages"`
	Configs             []ProfileConfig `yaml:"configurations"`
	PostInstallCommands []string        `yaml:"commands"`
	Repository          string          `yaml:"repository"`
	UninstallCommands   []string        `yaml:"uninstall"`
	Installed           bool
}

type ProfileConfig struct {
	Command CommandType `yaml:"command"`
	From    string      `yaml:"from"`
	To      string      `yaml:"to"`
}

type ProfilePackages struct {
	Install []string `yaml:"install"`
	Remove  []string `yaml:"remove"`
}

func (self Profile) InstallConfigFiles() (errs []error) {
	for _, config := range self.Configs {
		from, err := ExpandPath(config.From)
		if err != nil {
			errs = append(errs, err)
		}
		if _, err := os.Stat(from); !os.IsNotExist(err) {
			to, err := ExpandPath(config.To)
			if err != nil {
				errs = append(errs, err)
			}
			fmt.Println("To path [" + to + "] already exists, deleting...")
			err = terminal(`rm ` + to)
			if err != nil {
				errs = append(errs, err)
			}
			fmt.Println("Using [", config.Command.String(), "] on \"", from, "\" -> \"", to, "\"")
			err = terminal(config.Command.Execute() + ` "` + from + `" "` + to + `"`)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			errs = append(errs, errors.New("[Error] Failed to ["+config.Command.String()+"] config file ["+config.From+"]->["+config.To+"]"))
		}
	}
	return errs
}

func (self Profile) ExecutePostInstallCommands() (errs []error) {
	for _, cmd := range self.PostInstallCommands {
		// TODO: Should validate, have a limited set of commands accessible to the
		// provisioning script, and defined input for each allowed command
		// (whitelist)
		// This will avoid erroring on things like symbolic link existing
		if strings.Contains(cmd, "&&") {
			cmd = `/bin/sh -c "` + cmd + `"`
		}
		terminal(cmd)
	}
	return nil
}

func DefaultEnvironment() Environment {
	return Environment{
		Distribution: Distribution{
			OS:      Debian,
			Version: Version{Major: 9, Minor: 9, Patch: 0},
		},
		Profiles: []Profile{
			Profile{
				Category: "default",
				Type:     "developer",
				//Subtype:  "",
				Packages: ProfilePackages{
					Install: []string{"curl", "git", "vim"},
					Remove:  []string{"nano"},
				},
				Configs: []ProfileConfig{
					ProfileConfig{
						Command: Copy,
						From:    "profiles/default/dot.bashrc",
						To:      "~/.bashrc",
					},
					ProfileConfig{
						Command: Copy,
						From:    "profiles/default/dot.gitconfig",
						To:      "~/.gitconfig",
					},
				},
				PostInstallCommands: []string{"echo \"dot.config default profile install complete\""},
			},
		},
	}
}

func LoadEnvironment(path string) (env Environment, err error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return Environment{}, err
	}
	err = yaml.Unmarshal(yamlFile, &env)
	if err != nil {
		return Environment{}, err
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
