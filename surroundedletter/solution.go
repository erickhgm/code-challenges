package surroundedletter

import (
	"unicode"
	"unicode/utf8"
)

func Surrounded(text string) bool {
	if len(text) < 3 {
		return false
	}

	first, _ := utf8.DecodeRune([]byte{text[0]})
	last, _ := utf8.DecodeRune([]byte{text[len(text)-1]})
	// First and Last can not be letters.
	if unicode.IsLetter(first) || unicode.IsLetter(last) {
		return false
	}

	previous := first
	for _, v := range text[1:] {
		// Checking if we have two letters together.
		if unicode.IsLetter(v) && unicode.IsLetter(previous) {
			return false
		}
		previous = v
	}
	return true
}
