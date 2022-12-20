package decapitalizeconst

import (
	"strings"
)

func Decapitalize(s string) string {
	removeUndercore := func(s []string) []string {
		var ss []string
		for _, v := range s {
			if v != "" {
				ss = append(ss, strings.ToLower(v))
			}
		}
		return ss
	}
	parts := removeUndercore(strings.Split(s, "_"))

	var words []string
	for i, v := range parts {
		if i != 0 {
			v = strings.ToUpper(v[:1]) + v[1:]
		}
		words = append(words, v)
	}
	return strings.Join(words, "")
}
