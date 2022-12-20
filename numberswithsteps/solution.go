package numberswithsteps

func Steps(n, step int) []int {
	if n == 0 {
		return []int{}
	}
	var steps []int
	for i := n; i >= step; i = i - step {
		steps = append(steps, i)
	}
	return steps
}
