package pairwithaverage

func HasAverage(n []int, av float32) bool {
	if len(n) == 0 {
		return false
	}

	p1 := 0
	p2 := len(n) - 1

	for p1 != p2 {
		v1 := n[p1]
		v2 := n[p2]
		value := float32(v1+v2) / 2

		if value == av {
			return true
		} else if value > av {
			p2--
		} else if value < av {
			p1++
		}
	}
	return false
}
