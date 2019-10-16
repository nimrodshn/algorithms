package quicksort

// QuickSort is an implementation of the quicksort algorithm.
func QuickSort(input []int) []int {
	return sort(input, 0, len(input)-1)
}

func sort(input []int, p, r int) []int {
	if p < r {
		q := partition(input, p, r)
		sort(input, p, q-1)
		sort(input, q+1, r)
	}
	return input
}

func partition(input []int, p, r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if input[j] <= input[r] {
			i++
			k := input[i]
			input[i] = input[j]
			input[j] = k
		}
	}
	k := input[i+1]
	input[i+1] = input[r]
	input[r] = k
	return i + 1
}
