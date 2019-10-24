package bucketsort

import (
	"github.com/nimrodshn/algorithms/insertionsort"
)

// BucketSort is an implementation of the bucket sort algorithm.
func BucketSort(A []int, m int) []int {
	n := len(A)
	buckets := make([][]int, n)
	result := make([]int, 0)
	for _, x := range A {
		buckets[(n*x)/m] = append(buckets[(n*x)/m], x)
	}
	for i := range buckets {
		buckets[i] = insertionsort.InsertionSort(buckets[i])
		result = append(result, buckets[i]...)
	}
	return result
}
