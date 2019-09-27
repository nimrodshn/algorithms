package mergesort

// Merge merges two sorted sub arrays into one sorted array.
func Merge(left, right []int) []int {
	output := make([]int, len(right)+len(left))
	var i, j int
	for k := 0; k < len(left)+len(right); k++ {
		if right[i] < left[j] {
			output[k] = right[i]
			i++
			// If reached the end for the right pile
			// copy the rest of the left pile over.
			if i == len(right) {
				k++
				for ; j < len(left); j++ {
					output[k] = left[j]
					k++
				}
				break
			}
		} else {
			output[k] = left[j]
			j++
			// If reached the end for the left pile
			// copy the rest of the right pile over.
			if j == len(left) {
				k++
				for ; i < len(right); i++ {
					output[k] = right[i]
					k++
				}
				break
			}
		}
	}
	return output
}

func MergeSort(input []int) []int {
	if input == nil || len(input) < 2 {
		return input
	}
	k := len(input) / 2
	left := MergeSort(input[:k])
	right := MergeSort(input[k:])
	return Merge(left, right)
}
