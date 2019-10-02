package binaryaddition

// BinaryAddition adds two binary numbers represented as int slices of length n.
// An implementation of ex. 2.1-4 in Cormen's Introduction to Algorithms
func BinaryAddition(A, B []int, n int) []int {
	if n == 0 {
		return nil
	}
	carry := 0
	result := make([]int, n+1)
	for i := 0; i < n; i++ {
		result[i] = (A[i] + B[i] + carry) % 2
		carry = A[i] * B[i]
	}
	result[n] = carry
	return result
}
