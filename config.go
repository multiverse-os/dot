package main

import (
	"io/ioutil"
)

type Settings map[string]string

type ConfigurationFile struct {
	Application    string
	CustomSettings Settings
	Path           string
	Filename       string
}

func (self Configuration) String() string { return (self.Template + self.Settings) }
func (self Configuration) Path() string   { return (self.Path + self.Filename) }
func (self Configuration) Install() error { return ioutil.WriteFile(self.Path, self.String(), 0660) }

func Install(name string, settings Settings) (Config, error) {
	//switch {
	//case "neovim":
	//	return neovim(settings), nil
	//case "bash":
	//	return bash(settings), nil
	//case "git":
	//	return git(setting), nil
	//default:
	//	return nil, errors.New("configuration not supported")
	//}
}
