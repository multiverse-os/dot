package dot

type CommandType int

const (
	Copy CommandType = iota
	Link
	Remove
	Move
	MakeDirectory
)

// Aliases
const (
	Cp    = Copy
	Ln    = Link
	Rm    = Remove
	Mv    = Move
	MkDir = MakeDirectory
)

func (self CommandType) String() string {
	switch self {
	case Link:
		return "ln"
	case Remove:
		return "rm"
	case Move:
		return "mv"
	case MakeDirectory:
		return "mkdir"
	default: // CopyCommand
		return "cp"
	}
}

func (self CommandType) Execute() string {
	switch self {
	case Link:
		return "ln -s "
	case Remove:
		return "rm -rf "
	case Move:
		return "mv "
	case MakeDirectory:
		return "mkdir -p "
	default: // Copy
		return "cp -rf "
	}
}

func MarshalCommand(cmd string) CommandType {
	switch cmd {
	case Link.String():
		return Link
	case Remove.String():
		return Remove
	case Move.String():
		return Move
	case MakeDirectory.String():
		return MakeDirectory
	default: // CopyCommand.String()
		return Copy
	}
}
