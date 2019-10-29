package quicksort

import (
	"github.com/nimrodshn/algorithms/partition"
)

// QuickSort is an implementation of the quicksort algorithm.
func QuickSort(input []int) []int {
	return sort(input, 0, len(input)-1)
}

func sort(input []int, p, r int) []int {
	if p < r {
		q := partition.Lomuto(input, p, r)
		sort(input, p, q-1)
		sort(input, q+1, r)
	}
	return input
}
