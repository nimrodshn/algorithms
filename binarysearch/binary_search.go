package binarysearch

// BinarySearch finds an element x in a sorted arrray
// it returns the index of the element x in the array or -1,
// if no such element exist.
func BinarySearch(input []int, x int) int {
	r := len(input)
	l := 0
	for l < r-1 {
		if x == input[(r+l)/2] {
			return (r + l) / 2
		}
		if x > input[(r+l)/2] {
			l = (r + l) / 2
		} else {
			r = (r + l) / 2
		}
	}
	return -1
}
