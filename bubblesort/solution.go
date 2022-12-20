package bubblesort

func Sort(n []int) []int {
	var swapped bool
	for {
		swapped = false
		for i := 0; i < len(n)-1; i++ {
			left := n[i]
			right := n[i+1]
			if n[i] > n[i+1] {
				n[i] = right
				n[i+1] = left
				swapped = true
			}
		}
		if !swapped {
			return n
		}
	}
}
