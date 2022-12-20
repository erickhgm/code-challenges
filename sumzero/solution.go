package sumzero

func SumZero(n []int) []int {
	if len(n) == 0 {
		return []int{}
	}

	p1 := 0
	p2 := len(n) - 1

	for p1 != p2 {
		v1 := n[p1]
		v2 := n[p2]
		value := v1 + v2

		if value == 0 {
			return []int{v1, v2}
		} else if value > 0 {
			p2--
		} else if value < 0 {
			p1++
		}
	}
	return []int{}
}
