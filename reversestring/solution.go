package reversestring

func Reverse(s string) string {
	var reverse string
	for _, v := range s {
		reverse = string(v) + reverse
	}
	return reverse
}
