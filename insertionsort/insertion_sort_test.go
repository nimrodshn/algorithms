package insertionsort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
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
			expected: []int{1, 2, 3, 4, 6, 7, 9},
		},
	}
	for _, test := range tests {
		result := InsertionSort(test.input)
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
