package quicksort

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

//
//
//

func quickSort2(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition2(arr, start, end)
	quickSort2(arr, start, pivot-1)
	quickSort2(arr, pivot+1, end)
}

// [5, 6, 7, 2, 1, 0]
// [0, 6, 7, 2, 1, 5]
// [0, 2, 7, 6, 1, 5] => [0, 2, 1, 6, 7, 5] => [0, 2, 1, 5, 7, 6]
func partition2(arr []int, start, end int) int {
	pivotIdx := start
	pivotValue := arr[end]

	for i := start; i < end; i++ {
		if arr[i] < pivotValue {
			arr[i], arr[pivotIdx] = arr[pivotIdx], arr[i]
			pivotIdx++
		}
	}

	arr[pivotIdx], arr[end] = arr[end], arr[pivotIdx]
	return pivotIdx
}
