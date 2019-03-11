package dot

type ProfileType int

const (
	DefaultProfile ProfileType = iota
	DeveloperProfile
)

type DeveloperSubtype int

// TODO: Will need the ability to mix and match in the future
const (
	CLanguage DeveloperSubtype = iota // Default
	GoLanguage
	RustLanguage
	RubyLanguage
	PythonLanguage
)

type Profile struct {
	Type     ProfileType
	Language DeveloperSubtype
	Filename string
}

// TODO: Load from yaml goes here
