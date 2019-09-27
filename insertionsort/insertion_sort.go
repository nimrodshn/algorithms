package insertionsort

func InsertionSort(input []int) []int {
	if input == nil || len(input) <= 1 {
		return input
	}
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if j == 0 {
				continue
			}
			if input[j] < input[j-1] {
				k := input[j-1]
				input[j-1] = input[j]
				input[j] = k
			}
		}
	}
	return input
}
