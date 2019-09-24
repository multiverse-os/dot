package dot

import ()

type Settings map[string]string

type Config struct {
	Application string
	Settings    Settings
	Path        string
	Filename    string
	Template    string
}
