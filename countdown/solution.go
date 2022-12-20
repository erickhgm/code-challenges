package countdown

func Down(c int) []int {
	count := []int{}
	for i := c; i >= 0; i-- {
		count = append(count, i)
	}
	return count
}
