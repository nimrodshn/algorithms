package binaryaddition

import (
	"testing"
)

func TestBinaryAddition(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		len      int
		expected []int
	}{
		{
			A:        []int{1, 0, 1},
			B:        []int{1, 0, 0},
			len:      3,
			expected: []int{0, 1, 1, 0},
		},
		{
			A:        []int{1, 1, 1},
			B:        []int{1, 1, 1},
			len:      3,
			expected: []int{0, 1, 1, 1},
		},
	}
	for _, test := range tests {
		result := BinaryAddition(test.A, test.B, test.len)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Fatalf("Test Failed: expected %v, instead recieved: %v",
					test.expected, result)
			}
		}
	}
}
