package permutationpalindrome

func IsPalindromePermutation(text string) bool {
	letters := map[string]bool{}
	for _, v := range text {
		temp := string(v)

		if ok := letters[temp]; ok {
			delete(letters, temp)
		} else {
			letters[temp] = true
		}
	}
	return len(letters) <= 1
}
