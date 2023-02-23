package factorial

import "fmt"

func Factorial(n int64) int64 {
	fac := int64(1)
	for i := int64(1); i <= n; i++ {
		fac *= i
	}
	return fac
}

// FactorialRecursive calculates the Factorial recursivily.
// 3! => 3 * 2 * 1 => 6
// 0! = 1
// 1! = 1
func FactorialRecursive(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

// 1, 2, 3, 5
func PrintPermutations(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
		return
	}

	for i := 0; i < size; i++ {
		PrintPermutations(a, size-1)

		if size%2 != 0 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}
