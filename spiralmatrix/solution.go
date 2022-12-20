package spiralmatrix

func emptyMatrix(n int) [][]int {
	spiral := [][]int{}
	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < n; j++ {
			temp = append(temp, 0)
		}
		spiral = append(spiral, temp)
	}
	return spiral
}

func Generate(n int) [][]int {
	left, top := 0, 0
	right, botton := n-1, n-1
	spiral := emptyMatrix(n)
	value := 1
	// [
	//  [1, 2, 3],
	//  [8, 9, 4],
	//  [7, 6, 5]
	// ]
	for left <= right {

		for c := left; c <= right; c++ {
			spiral[top][c] = value
			value++
		}
		top++

		for r := top; r <= botton; r++ {
			spiral[r][right] = value
			value++
		}
		right--

		for c := right; c >= left; c-- {
			spiral[botton][c] = value
			value++
		}
		botton--

		for r := botton; r >= top; r-- {
			spiral[r][left] = value
			value++
		}
		left++
	}
	return spiral
}
