package binarysearch

func Search(elmts []string, elmt string) int {
	left := 0
	right := len(elmts) - 1

	for left <= right {
		mid := (right + left) / 2
		v := elmts[mid]

		if elmt == v {
			return mid
		}

		if elmt > v {
			left = mid + 1 // x greater than middle, so ignore left half
		} else {
			right = mid - 1 // x greater than middle, so ignore right half
		}
	}

	return -1
}

// 0 1 2 3 4 5 6 7 8 9
// A B C D E F G H I J

func binarySearch(items []int, item int) int {
	left := 0
	right := len(items) - 1

	for left <= right {
		mid := (left + right) / 2
		current := items[mid]

		if current == item {
			return mid
		}

		if item < current {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
