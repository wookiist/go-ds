package insertionsort

// InsertionSort sorts an integer slice with insertionsort-way
func InsertionSort(arr *[]int) {
	var tempVal int
	n := len(*arr)
	for i := 1; i < n; i++ {
		tempVal = (*arr)[i]
		for j := i - 1; j >= 0; j-- {
			if (*arr)[j] > tempVal {
				(*arr)[j+1], (*arr)[j] = (*arr)[j], (*arr)[j+1]
			} else {
				(*arr)[j+1] = tempVal
				break
			}
		}
	}
}
