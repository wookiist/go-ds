package selectionsort

// SelectionSort sorts an integer array with selectionsort-way
func SelectionSort(arr *[]int) {
	var min int
	n := len(*arr)
	for i := 0; i < n; i++ {
		min = i
		for j := i; j < n; j++ {
			if (*arr)[min] > (*arr)[j] {
				min = j
			}
		}
		(*arr)[i], (*arr)[min] = (*arr)[min], (*arr)[i]
	}
}
