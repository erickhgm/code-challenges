package digitfrequency

import "reflect"

func EqualsFrequency(a, b string) bool {
	aFreq := map[string]int{}
	bFreq := map[string]int{}

	for _, v := range a {
		k := string(v)
		if f, ok := aFreq[k]; ok {
			aFreq[k] = f + 1
		} else {
			aFreq[k] = 1
		}
	}
	for _, v := range b {
		k := string(v)
		if f, ok := bFreq[k]; ok {
			bFreq[k] = f + 1
		} else {
			bFreq[k] = 1
		}
	}
	return reflect.DeepEqual(aFreq, bFreq)
}
