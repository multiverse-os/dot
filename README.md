<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse: `dot.config` dot config manager and simple system provisioning
**URL** [multiverse-os.org](https://multiverse-os.org)

Using YAML (and JSON) to define development environments or general system
configurations, from simple dot configuration installations to basic system
provisioning. An environment is a one or more profiles; a profile is a a
collection of packages to install and remove, configuration files to install,
and post installation commands to run.

Application based provisioning using a YAML (or JSON) will be an additional
profile option. Each applicaiton configuration will have a package,
dependencies, configurations (supporting custom configuration values), and post
installation commands. 



*API is still under heavy development (this project is entirely experimental,
its a learning tool for a new developer).*


```
ln -s ~/.config/nvim/init.vim ~/.vimrc


## Persistent undo
mkdir ~/.config/nvim/undo
mkdir ~/.config/nvim/backup
mkdir ~/.config/nvim/tmp

echo "set undofile" >> ~/.vimrc
echo "undodir=~/.config/nvim/undo" >> ~/.vimrc
echo "backupdir=~/.config/nvim/backup" >> ~/.vimrc
echo "directory=~/.config/nvim/tmp" >> ~/.vimrc
```

