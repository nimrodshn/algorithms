package binarysearch

import "testing"

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		x        int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			x:        2,
			expected: 1,
		},
		{
			input:    []int{1, 2, 3},
			x:        10,
			expected: -1,
		},
		{
			input:    []int{1, 2, 3, 6, 7, 9},
			x:        7,
			expected: 4,
		},
		{
			input:    []int{1, 2, 3, 6, 7, 9},
			x:        5,
			expected: -1,
		},
	}
	for _, test := range tests {
		result := BinarySearch(test.input, test.x)
		if result != test.expected {
			t.Fatalf("Test Failed: expected %v, instead recieved: %v",
				test.expected, result)
		}
	}
}
