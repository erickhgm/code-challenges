package isanagram

import (
	"reflect"
	"strings"
	"unicode"
)

func countFrequency(s string) map[string]bool {
	letters := map[string]bool{}
	for _, v := range s {
		if !unicode.IsLetter(v) {
			continue
		}
		temp := strings.ToLower(string(v))
		if _, ok := letters[temp]; !ok {
			letters[temp] = true
		} else {
			delete(letters, temp)
		}
	}
	return letters
}

func IsAnagram(a string, b string) bool {
	am := countFrequency(a)
	bm := countFrequency(b)

	return reflect.DeepEqual(am, bm)
}
