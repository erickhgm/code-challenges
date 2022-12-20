package maxoccurrentchar

func MaxChar(s string) string {
	letters := map[string]int{}
	for _, v := range s {
		if _, ok := letters[string(v)]; ok {
			letters[string(v)] = letters[string(v)] + 1
		} else {
			letters[string(v)] = 1
		}
	}
	max := 0
	var letter string
	for k, v := range letters {
		if v > max {
			letter = k
			max = v
		}
	}
	return letter
}
