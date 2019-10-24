package selection

// MinMaxSelection returns the minimum and maximum of the input array
// in worst case 3*n/2 comparison.
func MinMaxSelection(A []int) (int, int) {
	min := A[0]
	max := A[0]
	var larger, smaller int
	i := 0
	j := len(A) - 1
	for i <= j {
		if A[i] < A[j] {
			larger = A[j]
			smaller = A[i]
		} else {
			larger = A[i]
			smaller = A[j]
		}
		if max < larger {
			max = larger
		}
		if min > smaller {
			min = smaller
		}
		i++
		j--
	}
	return max, min
}
