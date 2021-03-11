package stringutil

import (
	"strings"
	"unicode"
)

func CleanStr(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, str)
}
