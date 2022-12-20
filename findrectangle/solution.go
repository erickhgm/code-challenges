package findrectangle

func getBottomRight(matrix [][]int, i, j int) (int, int) {
	var value int

	rightColumn := j
	for rightColumn < len(matrix[i]) {
		value = matrix[i][rightColumn]
		if value == 0 {
			rightColumn++
		} else {
			rightColumn--
			break
		}
	}

	bottonRow := i
	for bottonRow < len(matrix) {
		value = matrix[bottonRow][rightColumn]
		if value == 0 {
			bottonRow++
		} else {
			bottonRow--
			break
		}
	}
	return bottonRow, rightColumn
}

func Coordinates(matrix [][]int) []int {
	coordinates := []int{0, 0, 0, 0}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			value := matrix[i][j]

			if value == 0 {
				coordinates[0] = i
				coordinates[1] = j
				coordinates[2], coordinates[3] = getBottomRight(matrix, i, j)
				return coordinates
			}
		}
	}
	return coordinates
}
