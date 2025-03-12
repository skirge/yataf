package general

import (
	"flag"
	_struct "github.com/dsecuredcom/yataf/pkg/struct"
	"strings"
)

func GetArguments() _struct.Config {

	File := flag.String("file", "", "File to analyze")
	Types := flag.String("types", "all", "Type of regex to use")
	CharacterLimit := flag.Int("char-limit", 1000, "Character limit for data extracted")
	Meg := flag.Bool("meg", false, "Meg compatible mode (skip request headers)")

	flag.Parse()

	TypesSplit := strings.Split(strings.TrimSpace(*Types), ",")

	return _struct.Config{
		File:      *File,
		Types:     TypesSplit,
		CharLimit: *CharacterLimit,
		Meg:	*Meg,
	}
}
