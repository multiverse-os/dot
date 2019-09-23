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
	MultiverseProfile
	DevelopmentProfile
)

func (self ProfileType) String() string {
	switch self {
	case DevelopmentProfile:
		return "developer"
	case MultiverseProfile:
		return "multiverse"
	default:
		return "default"
	}
}

// NOTE: Profiles contain configs, and environments contain profiles.
// This allows for mixing and matching configs, or collections of configs
// (profiles), to make up (environment) collections of profiles. An environment
// can be defined by a git repository, or files on the system.
type Environment struct {
	OS         string    `yaml:"os"`
	Version    Version   `yaml:"version"`
	Repository string    `yaml:"git"`
	Profiles   []Profile `yaml:"profiles"`
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
	Category            string                 `yaml:"category"`
	Type                string                 `yaml:"type"`
	Subtype             string                 `yaml:"subtype"`
	Packages            ProfilePackages        `yaml:"packages"`
	ConfigFiles         []ProfileConfigInstall `yaml:"configurations"`
	PostInstallCommands []string               `yaml:"commands"`
	Repository          string                 `yaml:"repository"`
	UninstallCommands   []string               `yaml:"uninstall"`
	Installed           bool
}

type ProfilePackages struct {
	Install []string `yaml:"install"`
	Remove  []string `yaml:"remove"`
}

type CommandType int

const (
	Copy CommandType = iota
	Link
	Remove
	Move
	MakeDirectory
)

// Aliases
const (
	Cp    = Copy
	Ln    = Link
	Rm    = Remove
	Mv    = Move
	MkDir = MakeDirectory
)

func (self CommandType) Name() string {
	switch self {
	case Link:
		return "ln"
	case Remove:
		return "rm"
	case Move:
		return "mv"
	case MakeDirectory:
		return "mkdir"
	default: // CopyCommand
		return "cp"
	}
}

func (self CommandType) Command() string {
	switch self {
	case Link:
		return "ln -s"
	case Remove:
		return "rm -rf"
	case Move:
		return "mv"
	case MakeDirectory:
		return "mkdir -p"
	default: // Copy
		return "cp -rf"
	}
}

func MarshalCommand(cmd string) CommandType {
	switch cmd {
	case Link.Name():
		return Link
	case Remove.Name():
		return Remove
	case Move.Name():
		return Move
	case MakeDirectory.Name():
		return MakeDirectory
	default: // CopyCommand.String()
		return Copy
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

func (self Profile) ConfigFiles() (errs []error) {
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

func DefaultEnvironment() Environment {
	return Environment{
		OS:      "debian",
		Version: Version{Major: 10, Minor: 0, Patch: 1}.String(),
		Profiles: []Profile{
			Profile{
				Type: "default",
				Packages: ProfilePackages{
					Install: []string{"curl", "git", "vim"},
					Remove:  []string{"nano"},
				},
				ConfigFiles: []ProfileConfigInstall{
					ProfileConfigInstall{
						Command: "cp",
						From:    "dot.bashrc",
						To:      "~/.bashrc",
					},
					ProfileConfigInstall{
						Command: CopyCommand,
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
