package capitalizesentence

import "strings"

const Sep = " "

func Cap(s string) string {
	words := strings.Split(s, Sep)
	for i, v := range words {
		words[i] = strings.ToUpper(v[0:1]) + v[1:]
	}
	return strings.Join(words, Sep)
}
