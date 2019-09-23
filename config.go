package dot

import (
	"io/ioutil"

	config "github.com/multiverse-os/dot/config"
)

type Settings map[string]string

type ConfigFile struct {
	Application    string
	CustomSettings Settings
	Path           string
	Filename       string
	Template       string
}

// TODO: Modify specific line in configuration
//       Replace line 75 with " "

// TODO: Substring replace; full file; specific line

// TODO: Render template by taking settings and putting them inline into the
// template

// TODO: Using a map insert, replace, find/replace, delete LINE

//func (self ConfigFile) String() string   { return self.Template }
func (self ConfigFile) FullPath() string { return (self.Path + self.Filename) }
func (self ConfigFile) Install() error {
	return ioutil.WriteFile(self.FullPath(), []byte(self.Template), 0660)
}

func Install(name string, settings Settings) (ConfigFile, error) {
	switch name {
	case "neovim", "vim":
		return config.Neovim(settings), nil
	case "bash":
		return config.Bash(settings), nil
	case "git":
		return config.Git(setting), nil
	default:
		return nil, errors.New("configuration not supported")
	}
	return ConfigFile{}, nil
}
