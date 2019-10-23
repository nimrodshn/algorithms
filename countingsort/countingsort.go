package countingsort

// CountingSort is an implementation of the counting sort.
func CountingSort(A []int, k int) []int {
	histogram := make([]int, k)
	result := make([]int, len(A))

	for _, x := range A {
		histogram[x] = histogram[x] + 1
	}
	for i := 1; i < len(histogram); i++ {
		histogram[i] += histogram[i-1]
	}
	for i := range A {
		result[histogram[A[i]]-1] = A[i]
		histogram[A[i]] = histogram[A[i]] - 1
	}
	return result
}
