package dot

type Language int

const (
	C Language = iota // Default
	Go
	Rust
	Ruby
	Python
)

func (self Language) String() string {
	switch self {
	case Go:
		return "Go"
	case Rust:
		return "Rust"
	case Ruby:
		return "Ruby"
	case Python:
		return "Python"
	default:
		return "C"
	}
}
