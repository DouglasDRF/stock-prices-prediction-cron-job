package stringutil

import (
	"strings"
	"time"
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

func GetCurrentTimeStr() string {
	return time.Now().String() + " "
}
