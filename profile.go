package dot

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

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
	OS         string    `yaml:"os"`
	Repository string    `yaml:"git"`
	Version    string    `yaml:"version"`
	Profiles   []Profile `yaml:"profiles"`
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

type CommandType int

const (
	CopyCommand CommandType = iota
	LinkCommand
)

func (self CommandType) Name() string {
	switch self {
	case LinkCommand:
		return "link"
	default: // CopyCommand
		return "copy"
	}
}

func (self CommandType) StringWithFlags() string {
	switch self {
	case LinkCommand:
		return "ln -s"
	default: // CopyCommand
		return "cp -rf"
	}
}

func MarshalCommand(cmd string) CommandType {
	switch cmd {
	case LinkCommand.Name():
		return LinkCommand
	default: // CopyCommand.String()
		return CopyCommand
	}
}

func ExpandPath(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	} else if path[:1] == "~" {
		currentUser, err := user.Current()
		if err != nil {
			return path, err
		}
		return filepath.Join(currentUser.HomeDir, path[1:]), nil
	} else if path[:1] != "/" {
		wd, err := os.Getwd()
		if err != nil {
			return path, err
		}
		return fmt.Sprintf(wd + "/" + path), nil
	} else {
		return path, nil
	}
}

func (self Profile) CopyOrLinkConfigFiles() (errs []error) {
	for _, configFile := range self.ConfigFiles {
		from, err := ExpandPath(configFile.From)
		if err != nil {
			errs = append(errs, err)
		}
		if _, err := os.Stat(from); !os.IsNotExist(err) {
			cmd := MarshalCommand(configFile.Command)
			to, err := ExpandPath(configFile.To)
			if err != nil {
				errs = append(errs, err)
			}
			fmt.Println("To path [" + to + "] already exists, deleting...")
			err = terminal(`rm ` + to)
			if err != nil {
				errs = append(errs, err)
			}
			fmt.Println("Using [", cmd.Name(), "] on \"", from, "\" -> \"", to, "\"")
			err = terminal(cmd.StringWithFlags() + ` "` + from + `" "` + to + `"`)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			errs = append(errs, errors.New("[Error] Failed to ["+MarshalCommand(configFile.Command).StringWithFlags()+"] config file ["+configFile.From+"]->["+configFile.To+"]"))
		}
	}
	return errs
}

type ProfileConfigInstall struct {
	Command string `yaml:"command"`
	From    string `yaml:"from"`
	To      string `yaml:"to"`
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
						Command: "copy",
						From:    "dot.gitconfig",
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
