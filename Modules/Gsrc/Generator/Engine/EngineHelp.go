package Engine

import (
	"strings"
	"unicode/utf8"
)

const (
	s = 60
	e = 62
)

type HTMLData struct {
	Startb int
	Endb   int
}

func (dt *HTMLData) StripHTML(str string) (NewBuild string) {
	var builder strings.Builder
	dt.Endb = 0
	dt.Startb = 0
	builder.Grow(len(str) + utf8.UTFMax)
	in := false // True if we are inside an HTML tag.
	for i, char := range str {
		if (i+1) == len(str) && dt.Endb >= dt.Startb {
			builder.WriteString(str[dt.Endb:])
		}
		if char != s && char != e {
			continue
		}
		if char == s {
			if !in {
				dt.Startb = i
			}
			in = true
			builder.WriteString(str[dt.Endb:dt.Startb])
			continue
		}
		in = false
		dt.Endb = i + 1
	}
	str = builder.String()
	NewBuild = str
	return NewBuild
}
