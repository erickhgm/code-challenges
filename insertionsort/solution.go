package insertionsort

func Sort(in []int) []int {
	for cIndex := 1; cIndex < len(in); cIndex++ {
		cValue := in[cIndex]

		pIndex := cIndex - 1
		pValue := in[pIndex]

		//fmt.Println(pValue, " > ", cValue)
		for pValue > cValue {
			in[pIndex+1] = pValue

			pIndex--
			if pIndex < 0 {
				break
			}
			pValue = in[pIndex]
			//fmt.Println(pValue, " > ", cValue)
		}
		in[pIndex+1] = cValue
	}
	return in
}
