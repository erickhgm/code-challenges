package template

func IsPalindrome(s string) bool {
	reverse := func(s string) string {
		var result string
		for _, v := range s {
			result = string(v) + result
		}
		return result
	}

	return s == reverse(s)
}
