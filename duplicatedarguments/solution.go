package duplicatedarguments

func Dup(args []string) []string {
	freq := map[string]int{}
	for _, v := range args {
		k := string(v)
		if f, ok := freq[k]; ok {
			freq[k] = f + 1
		} else {
			freq[k] = 1
		}
	}
	dup := []string{}
	for k, v := range freq {
		if v > 1 {
			dup = append(dup, k)
		}
	}
	return dup
}
