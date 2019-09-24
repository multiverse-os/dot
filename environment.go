package dot

// NOTE: Profiles contain configs, and environments contain profiles.
// This allows for mixing and matching configs, or collections of configs
// (profiles), to make up (environment) collections of profiles. An environment
// can be defined by a git repository, or files on the system.
type Environment struct {
	Distribution Distribution `yaml:"distribution"`
	Version      Version      `yaml:"version"`
	Repository   string       `yaml:"git"`
	Profiles     []Profile    `yaml:"profiles"`
}
