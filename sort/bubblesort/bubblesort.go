package bubblesort

// BubbleSort sorts an integer array with bubblesort-way
func BubbleSort(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
