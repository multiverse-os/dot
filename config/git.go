package config

type Git struct {
	application string
	filename string
	path string
	template string
	settings Settings
}

func Init(settings Settings) (Git, error) {
	return ConfigFile{
		application: "git",
		filename: ".gitconfig",
		path: "~/",
		settings: settings,
		template: gitTemplate(),
	}
}

func gitTemplate() string {
		return `
[user]
	email = you@example.com
	name = Your Name

[color]
  diff = auto
  status = auto
  branch = auto
  interactive = auto
  ui = true
  pager = true

[color "status"]
  added = green
  changed = red bold
  untracked = magenta bold

[color "branch"]
  remote = yellow
`}
}
