package linearsearch

import "strings"

func IndexOf(elmts []string, elmt string) int {
	for i, e := range elmts {
		if strings.EqualFold(e, elmt) {
			return i
		}
	}
	return -1
}
