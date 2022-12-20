package squereequals

import "reflect"

func Square(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aFreq := map[int]int{}
	for _, v := range a {
		square := v * v
		if f, ok := aFreq[square]; ok {
			aFreq[square] = f + 1
		} else {
			aFreq[square] = 1
		}
	}
	bFreq := map[int]int{}
	for _, v := range b {
		if f, ok := bFreq[v]; ok {
			bFreq[v] = f + 1
		} else {
			bFreq[v] = 1
		}
	}
	return reflect.DeepEqual(aFreq, bFreq)
}
