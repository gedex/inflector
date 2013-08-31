package inflector

import (
	"strings"
)

// A rule
type Rule int

const (
	Plural = iota
	Singular
	Transliteration
)
