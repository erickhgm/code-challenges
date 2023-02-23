package template

// func MergeSort(elmts []int) []int {
// 	if len(elmts) < 2 {
// 		return elmts
// 	}
// 	middleIdx := len(elmts) / 2
// 	return merge(MergeSort(elmts[:middleIdx]), MergeSort(elmts[middleIdx:]))
// }

// // left=[1, 3, 5] right=[2, 4, 6]
// func merge(left, right []int) []int {
// 	var merged []int
// 	var leftIdx int
// 	var rightIdx int

// 	for leftIdx < len(left) && rightIdx < len(right) {
// 		if left[leftIdx] < right[rightIdx] {
// 			merged = append(merged, left[leftIdx])
// 			leftIdx++
// 		} else {
// 			merged = append(merged, right[rightIdx])
// 			rightIdx++
// 		}
// 	}
// 	lastElmt := append(left[leftIdx:], right[rightIdx:]...)
// 	merged = append(merged, lastElmt...)
// 	return merged
// }

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items) / 2
	return merge(MergeSort(items[:mid]), MergeSort(items[mid:]))
}

func merge(left, right []int) []int {
	var merged []int
	var leftIdx int
	var rightIdx int

	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] < right[rightIdx] {
			merged = append(merged, left[leftIdx])
			leftIdx++
		} else {
			merged = append(merged, right[rightIdx])
			rightIdx++
		}
	}

	last := append(left[leftIdx:], right[rightIdx:]...)
	merged = append(merged, last...)
	return merged
}
