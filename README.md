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

The dot configurations repository allows the user to define the repository
source the configurations. By default `dot-config` will look for the
`base-files`, or `profiles/base-files`. The `base-files` path is expected to
contain folders corresponding to the `type`. Each of these `type` named folders
should be all the required files relative to the `~/` path.


````
os: "debian"
git: "github.com/multiverse-os/dot-configs"
version: "9.7"
profiles:
- type: "development"
  subtype: "golang"
  packages:
    install:
    - "build-essential"
    - "cmake"
    - "curl"
    - "git"
    - "golang"
    - "neovim"
    - "vim-gocomplete"
    remove:
    - "nano"
  configurations:
  - from: "dot.config/nvim/init.vim"
    to: "~/.config/nvim/init.vim"
  - from: "dot.bashrc"
    to: "~/.bashrc"
  commands:
  - "cd ~/ && ln -s ~/.config/nvim/init.vim ~/.vimrc"
````

For example, `base-files/development/` will contain `../development/dot.gitconfig`
that will be installed in the local user's home folder. This enables users to
store their own dot config files in a git repository, and store config files for
several machines or break up config files into modules that can be mixed and
matched for maximum customability with the simplest interface possible. 

The subtype allows further organization by having a subfolder within the type
folder, for example the `base-files` for the above type "development" and
subtype "golang" would be `~/base-files/development/golang/*`.


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
````

