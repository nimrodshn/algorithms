package heap

import "testing"

func TestBuildMaxHeap(t *testing.T) {
	tests := []struct {
		input    []int
		expected *MaxHeap
	}{
		{
			input: []int{1, 2, 3},
			expected: &MaxHeap{
				heap: []int{3, 2, 1},
			},
		},
		{
			input: []int{1, 2, 3, 6, 7, 9},
			expected: &MaxHeap{
				heap: []int{9, 7, 3, 6, 2, 1},
			},
		},
	}
	for _, test := range tests {
		result := BuildMaxHeap(test.input)
		if result.Size() != test.expected.Size() {
			t.Fatalf("Test Failed: expected %v, instead recieved: %v",
				test.expected, result)
		}
		for i := range result.Array() {
			if result.Get(i) != test.expected.Get(i) {
				t.Fatalf("Test Failed: expected %v, instead recieved: %v",
					test.expected, result)
			}
		}
	}
}

func TestHeapSort(t *testing.T) {
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
			input:    []int{1, 2, 3, 7, 6, 9, 4},
			expected: []int{1, 2, 3, 4, 6, 7, 9},
		},
		{
			input:    []int{31, 41, 59, 26, 41, 58},
			expected: []int{26, 31, 41, 41, 58, 59},
		},
	}
	for _, test := range tests {
		result := Sort(test.input)
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
