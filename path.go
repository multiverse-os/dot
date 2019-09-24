package dot

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

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
