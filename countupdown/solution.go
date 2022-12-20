package countupdown

func UpAndDown(n int) []int {
	var upAndDown []int

	for i := 0; i < n; i++ {
		upAndDown = append(upAndDown, i)
	}

	upAndDown = append(upAndDown, n)

	for i := n - 1; i >= 0; i-- {
		upAndDown = append(upAndDown, i)
	}
	return upAndDown
}
