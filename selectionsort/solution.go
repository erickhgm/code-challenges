package selectionsort

func Sort(n []int) []int {
	for i := 0; i < len(n); i++ {
		minIndex := i

		for j := i + 1; j < len(n); j++ {
			if n[j] < n[minIndex] {
				minIndex = j
			}
		}
		temp := n[minIndex]
		n[minIndex] = n[i]
		n[i] = temp
	}
	return n
}
