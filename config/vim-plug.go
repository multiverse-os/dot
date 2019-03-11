package dot

import (
	"fmt"
)

// TODO: There is a go project that handles vim plugin managers, probably work merging that instead of doing this

func InstallVimPlug() string {
	fmt.Println("Installing Vim Plug for neovim")
	return "curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
}
