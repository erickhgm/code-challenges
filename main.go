package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// TODO
	//input := [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3", "b4"}, {"c1"}, {"d1", "d2"}, {"e1"}}
	//output := []string{"a1", "b1", "c1", "d1", "e1", "a2", "b2", "d2", "a3", "b3", "b4"}

	str := "Hello apple pie"

	rep := map[string]int{}

	for _, s := range strings.Split(str, " ") {
		temp := map[string]int{}
		for _, v := range s {
			if _, ok := temp[string(v)]; ok {
				temp[string(v)] = temp[string(v)] + 1
			} else {
				temp[string(v)] = 1
			}
		}
		count := 0
		for _, v := range temp {
			if v > 1 {
				count++
			}
		}
		rep[s] = count
	}
	greater := 1
	for _, v := range rep {
		if v > greater {
			greater = v
		}
	}
	words := []string{}
	for k, v := range rep {
		if v == greater {
			words = append(words, k)
		}
	}
	var word string
	var end bool
	for _, s := range strings.Split(str, " ") {
		if end {
			break
		}
		for _, v := range words {
			if v == s {
				word = v
				end = true
				break
			}
		}
	}
	if len(word) == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(word)
	}

	// arr := []int{5, 2, 4, 6}
	// //out := []int{2,3,4}

	// window := arr[0]
	// numbers := arr[1:]

	// var result []int
	// // result = append(result, numbers[0])
	// // result = append(result, (numbers[0]+numbers[1])/2)
	// // result = append(result, (numbers[0]+numbers[1]+numbers[2])/3)

	// win := 1
	// var value []int
	// for win <= window && win < len(arr) {
	// 	temp := numbers[0:win]
	// 	result = append(result, getMedian(temp))
	// 	win++
	// }

	// for i := 1; i < len(numbers)-2; i++ {
	// 	win = i
	// 	value = []int{}
	// 	for win <= window {
	// 		value = append(value, numbers[win])
	// 		win++
	// 	}
	// 	result = append(result, getMedian(value))
	// 	window++
	// }

	// var output string
	// for i, v := range result {
	// 	if i == len(result)-1 {
	// 		output += fmt.Sprintf("%v", v)
	// 	} else {
	// 		output += fmt.Sprintf("%v,", v)
	// 	}
	// }
	// fmt.Println(output)
}

func getMedian(arr []int) int {
	if len(arr)%2 == 1 {
		index := len(arr) / 2
		sort.Ints(arr)
		return arr[index]
	} else {
		sort.Ints(arr)
		temp1 := arr[0]
		temp2 := arr[len(arr)-1]
		return (temp1 + temp2) / 2
	}
}
