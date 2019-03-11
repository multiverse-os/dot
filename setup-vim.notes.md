


ln -s ~/.config/nvim/init.vim ~/.vimrc


## Persistent undo
mkdir ~/.config/nvim/undo
mkdir ~/.config/nvim/backup
mkdir ~/.config/nvim/tmp

echo "set undofile" >> ~/.vimrc
echo "undodir=~/.config/nvim/undo" >> ~/.vimrc
echo "backupdir=~/.config/nvim/backup" >> ~/.vimrc
echo "directory=~/.config/nvim/tmp" >> ~/.vimrc


