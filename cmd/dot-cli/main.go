type dot

import (
	"fmt"

	dot "github.com/lostinblue/dot-config"
)

func main() {
	text.PrintBanner(
		(text.Light("[DOT:") + text.White("User Settings Manager") + text.Light("]")),
		(text.Light(" v") + text.Bold("0.1.0"),
	)
	dot.installProfile(dot.Config))
}

