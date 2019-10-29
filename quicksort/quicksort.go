package quicksort

import (
	"math/rand"
	"time"
)

// QuickSort is an implementation of the quicksort algorithm.
func QuickSort(input []int) []int {
	return sort(input, 0, len(input)-1)
}

func sort(input []int, p, r int) []int {
	if p < r {
		q := Partition(input, p, r)
		sort(input, p, q-1)
		sort(input, q+1, r)
	}
	return input
}

// Partition is an implementation of the Lomuto partition scheme
func Partition(input []int, p, r int) int {
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

// RandomPartition partitions the array according to a randomly picked element
// instead of the last element in the array.
func RandomPartition(input []int, p, r int) int {
	rand.Seed(time.Now().UnixNano())
	n := len(input)
	// pick a pivot at random and swap it with the last element.
	pivotIdx := rand.Intn(n)
	k := input[pivotIdx]
	input[pivotIdx] = input[r]
	input[r] = k
	return Partition(input, p, r)
}
