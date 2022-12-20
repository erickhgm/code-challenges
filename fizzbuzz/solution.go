package fizzbuzz

import "fmt"

func FizzBuzz(n int) []string {
	var elmts []string
	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			elmts = append(elmts, "FizzBuzz")
		} else if i%5 == 0 {
			elmts = append(elmts, "Buzz")
		} else if i%3 == 0 {
			elmts = append(elmts, "Fizz")
		} else {
			elmts = append(elmts, fmt.Sprintf("%v", i))
		}
	}
	return elmts
}
