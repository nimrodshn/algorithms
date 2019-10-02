package insertionsort

import (
	"testing"
)

func TestNonIncreasingInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			input:    []int{3, 2, 1},
			expected: []int{3, 2, 1},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    nil,
			expected: nil,
		},
		{
			input:    []int{1, 2, 3, 7, 6, 9, 4},
			expected: []int{9, 7, 6, 4, 3, 2, 1},
		},
	}
	for _, test := range tests {
		result := NonIncreasingInsertionSort(test.input)
		if len(result) != len(test.expected) {
			t.Fatalf("Test Failed: expected %v, instead recieved: %v",
				test.expected, result)
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Fatalf("Test Failed: expected %v, instead recieved: %v",
					test.expected, result)
			}
		}
	}
}
