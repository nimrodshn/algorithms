package insertionsort

// NonIncreasingInsertionSort an implementation of ex. 2.1-2 in Cormens Introduction to Algorithms.
func NonIncreasingInsertionSort(input []int) []int {
	if input == nil || len(input) <= 1 {
		return input
	}
	for i := 0; i < len(input); i++ {
		for j := i; j > 0; j-- {
			if input[j] > input[j-1] {
				k := input[j-1]
				input[j-1] = input[j]
				input[j] = k
			} else {
				break
			}
		}
	}
	return input
}
