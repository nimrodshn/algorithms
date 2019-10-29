package selection

import (
	"testing"
)

func TestPartitionSelect(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			k:        1,
			expected: 2,
		},
		{
			input:    []int{1, 2, 3, 7, 6, 9, 4},
			k:        3,
			expected: 4,
		},
		{
			input:    []int{31, 59, 26, 41, 58},
			k:        3,
			expected: 58,
		},
	}
	for _, test := range tests {
		result := PartitionSelect(test.input, test.k)
		if result != test.expected {
			t.Fatalf("Test Failed: expected %v, instead recieved: %v",
				test.expected, result)
		}

	}
}
