package tolerantpalindrome

func reverse(text string) string {
	var reversed string
	for _, v := range text {
		reversed = string(v) + reversed
	}
	return reversed
}

// abb#a
// a#bba
func IsPalindrome(text string) bool {
	reversed := reverse(text)
	if reversed == text {
		return true
	}

	for i := range text {
		temp1 := text[0:i] + text[i+1:]
		temp2 := reversed[0:len(text)-(i+1)] + reversed[len(text)-i:]
		if temp1 == temp2 {
			return true
		}
	}
	return false
}
