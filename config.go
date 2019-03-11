package dot

import (
	"io/ioutil"
)

type Settings map[string]string

type ConfigFile struct {
	Application    string
	CustomSettings Settings
	Path           string
	Filename       string
	Template       string
}

// TODO: Render template by taking settings and putting them inline into the
// template
//func (self ConfigFile) String() string   { return self.Template }
func (self ConfigFile) FullPath() string { return (self.Path + self.Filename) }
func (self ConfigFile) Install() error {
	return ioutil.WriteFile(self.FullPath(), []byte(self.Template), 0660)
}

func Install(name string, settings Settings) (ConfigFile, error) {
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
	return ConfigFile{}, nil
}
