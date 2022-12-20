package main

import "fmt"

func main() {
	// TODO
	input := [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3", "b4"}, {"c1"}, {"d1", "d2"}, {"e1"}}
	output := []string{"a1", "b1", "c1", "d1", "e1", "a2", "b2", "d2", "a3", "b3", "b4"}

	fmt.Println(input, output)
}
