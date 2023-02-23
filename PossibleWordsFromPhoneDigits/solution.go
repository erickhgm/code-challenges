package template

import "fmt"

var letters = map[int][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

// func MyFunc(numbers []int) {
// 	var combinations [][]string
// 	for _, v := range numbers {
// 		combinations = append(combinations, letters[v])
// 	}
// 	var combination []string
// 	iterate(combination, combinations)
// }

// func iterate(combination []string, numbers [][]string) {
// 	if len(numbers) == 0 {
// 		fmt.Println(combination)
// 		return
// 	}

// 	n, numbers := numbers[0], numbers[1:]

// 	for i := 0; i < len(n); i++ {

// 		copyCombination := append([]string{}, combination...)
// 		copyCombination = append(copyCombination, n[i])

// 		iterate(copyCombination, numbers)
// 	}
// }

func MyFunc(numbers []int) {
	var combinations [][]string
	for _, v := range numbers {
		combinations = append(combinations, letters[v])
	}

	var result []string
	iterate(combinations, result)
}

func iterate(combinations [][]string, result []string) {
	if len(combinations) == 0 {
		fmt.Println(result)
		return
	}

	comb, combinations := combinations[0], combinations[1:]

	for _, v := range comb {

		fmt.Println(v)

		result := append([]string{}, result...)
		result = append(result, v)

		iterate(combinations, result)
	}
}
