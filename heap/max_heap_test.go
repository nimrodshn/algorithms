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
		for i := range result.heap {
			if result.Get(i) != test.expected.Get(i) {
				t.Fatalf("Test Failed: expected %v, instead recieved: %v",
					test.expected, result)
			}
		}
	}
}
