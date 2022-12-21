package fibonacci

func Calc(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	fib := []int{0, 1}

	p2 := 0
	p1 := 1
	for i := 2; i <= n; i++ {
		value := p1 + p2
		p2 = p1
		p1 = value
		fib = append(fib, value)
	}
	return fib
}
