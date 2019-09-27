package mergesort

import (
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		left     []int
		right    []int
		expected []int
	}{
		{
			left:     []int{1, 4, 6},
			right:    []int{2, 3, 5},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			left:     []int{1},
			right:    []int{5},
			expected: []int{1, 5},
		},
	}
	for _, test := range tests {
		result := Merge(test.left, test.right)
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

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 4, 6, 2, 3, 5},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input:    []int{1, 5},
			expected: []int{1, 5},
		},
		{
			input:    []int{1, 5, 4},
			expected: []int{1, 4, 5},
		},
		{
			input:    []int{1, 2, 3, 7, 6, 9, 4},
			expected: []int{1, 2, 3, 4, 6, 7, 9},
		},
	}
	for _, test := range tests {
		result := MergeSort(test.input)
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
