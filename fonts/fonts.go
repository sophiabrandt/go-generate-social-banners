package fonts

import (
	_ "embed"
)

//go:embed BioRhyme-Bold.ttf
var TitleFont []byte

//go:embed FiraCode-Regular.ttf
var DefaultFont []byte
