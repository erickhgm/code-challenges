package template

import (
	"fmt"
	"strconv"
)

func IsPalindrome(s string) bool {
	reverse := func(s string) string {
		var result string
		for _, v := range s {
			result = string(v) + result
		}
		return result
	}

	return s == reverse(s)
}

func NextPalindrome(n int) int {

	var next int
	nstr := Number(n).ToString()
	length := len(nstr)
	isOdd := length%2 != 0

	if isOdd {
		// 123
		midIdx := length / 2
		right := reverse(nstr[:midIdx])

		temp := nstr[:midIdx+1] + right
		next = String(temp).ToInt()

		if next < n {
			nstr = Number(next).ToString()
			mid := nstr[midIdx]
			midInt := String(mid).ToInt()
			if midInt == 9 {
				left := nstr[:midIdx+1]
				leftInt := String(left).ToInt()
				leftInt++
				left = Number(leftInt).ToString()
				right := reverse(left[:midIdx])
				temp = Number(leftInt).ToString() + right
			} else {
				midInt++
				temp = nstr[:midIdx] + Number(midInt).ToString() + right
			}
			next = String(temp).ToInt()
		}

	} else {
		// 1234
		midIdx := length / 2
		right := reverse(nstr[:midIdx])

		temp := nstr[:midIdx] + right
		next = String(temp).ToInt()

		if next < n {
			nstr = Number(next).ToString()
			left := nstr[:midIdx]
			leftInt := String(left).ToInt()
			leftInt++

			left = Number(leftInt).ToString()
			right := reverse(left)
			temp = left + right
			next = String(temp).ToInt()
		}
	}

	return next
}

type Number int

func (n Number) ToString() string {
	return fmt.Sprintf("%v", n)
}

type String string

func (s String) ToInt() int {
	n, _ := strconv.Atoi(string(s))
	return n
}

func reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
