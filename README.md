<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse OS: `dot(config)` dot configuration file manager and simple system provisioning
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
git: "github.com/multiverse-os/dot-files"
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
  - command: "copy"
    from: "~/Go/src/github.com/multiverse-os/dot-config/profiles/dev.golang.yml"
    to: "~/test.yml"
  - command: "link"
    from: "~/Go/src/github.com/multiverse-os/dot-config/profiles/dev.ruby.yml"
    to: "~/test2.yml"
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

### Usage
Dot-config functions both as a library, and as a command-line tool named `dot`,
which can be used with custom profiles or by simply leveraging the default
built-in pseudoanonymous configurations such as configuring gitconfig with
`user@email.com`.

````
dot install default      # Default basic dot config file enhancements
dot install development  # Default built-in development configuration
````

Additionally `dot` includes by default a development environment configuration
that sets up a modern `neovim` environment using `vim-plug` and a collection of
popular vim plugins to enhance use for development for Go language. 

The standard usage is through specifying configuration through YAML (and JSON)
environment configurations which use one or more profile to setup an
environment. This configuration can even include a git repository, to use as the
source of dot configuration files. Using or not using `type` and `subtype`
indicuate to the `dot` command if only a single set of dot config files are
located in the repository, or they are organized into categories, or further
organized into categories and subcategories.

````
dot install dev.golang.yml
dot install https://remote.com/path/to.yml  # not yet implemented
````
-------------------------------------------------------------------------------
### Development
*API is still under heavy development (this project is entirely experimental,
its a learning tool for a new developer).*

#### Neovim configuration notes
Configuration of neovim is a common routine for developers who use `vim` but
additionally is complex enough to require the use of `exepct` like subprocess
automation tools in order to execute commands within `vim` such as `vim-plug`'s
`:PlugInstall` command from the command mode. This complexity makes it a great
example to build around, because if the process of downloading, installing
configuration files, setting up a plugin manager, then finally downloading
plugins can be completed; then almost any other dot configuration or applicaiton
setup can be accomplished. 


```
ln -s ~/.config/nvim/init.vim ~/.vimrc


#### Persistent undo
mkdir ~/.config/nvim/undo
mkdir ~/.config/nvim/backup
mkdir ~/.config/nvim/tmp

echo "set undofile" >> ~/.vimrc
echo "undodir=~/.config/nvim/undo" >> ~/.vimrc
echo "backupdir=~/.config/nvim/backup" >> ~/.vimrc
echo "directory=~/.config/nvim/tmp" >> ~/.vimrc
````

