package repeatedchar

func Repeated(s string) bool {
	letters := map[string]int{}
	for _, v := range s {
		if _, ok := letters[string(v)]; ok {
			letters[string(v)] = letters[string(v)] + 1
		} else {
			letters[string(v)] = 1
		}
	}
	for _, l := range letters {
		if l > 1 {
			return true
		}
	}
	return false
}
