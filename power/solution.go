package power

func Pow(base, exponent int) int {
	if exponent == 1 {
		return 1
	}
	result := base
	for i := 1; i < exponent; i++ {
		result = result * base
	}
	return result
}
