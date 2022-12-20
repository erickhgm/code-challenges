package getoddnumbers

func Odd(n []int) []int {
	odds := []int{}
	for _, v := range n {
		rest := v % 2
		if rest == 1 {
			odds = append(odds, v)
		}
	}
	return odds
}
