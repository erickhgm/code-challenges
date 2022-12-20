package factorial

func Factorial(n int) int {
	fac := 1
	for i := 1; i <= n; i++ {
		fac *= i
	}
	return fac
}
