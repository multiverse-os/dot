package config

import (
	"errors"
	"io/ioutil"
)

type Settings map[string]string
type ConfigFiles []ConfigFiles

type Config interface {
	InstallDependencies() error
	RemoveDependencies() error
	Install() error
	Uninstall() error
	Installed() bool
}

type ConfigFile struct {
	application string
	settings    Settings
	path        string
	filename    string
}

func (self ConfigFile) String() string { return (self.template + self.settings) }
func (self ConfigFile) Path() string   { return (self.path + self.filename) }
func (self ConfigFile) Install() error { return ioutil.WriteFile(self.Path, self.String(), 0660) }

func Install(name string, settings Settings) (Config, error) {
	switch {
	case "neovim":
		return neovim(settings), nil
	case "bash":
		return bash(settings), nil
	case "git":
		return git(setting), nil
	default:
		return nil, errors.New("configuration not supported")
	}
}
