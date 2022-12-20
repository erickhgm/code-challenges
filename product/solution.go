package product

func Product(n []int) int {
	var prod func(a int, b []int) int
	prod = func(a int, b []int) int {
		if len(b) == 0 {
			return a
		}
		return prod(a*b[0], b[1:])
	}
	return prod(1, n)
}
