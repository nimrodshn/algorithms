package countingsort

import "testing"

func TestCountingSort(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			k:        4,
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{3, 2, 1},
			k:        4,
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{1, 2, 3, 7, 6, 9, 4},
			k:        10,
			expected: []int{1, 2, 3, 4, 6, 7, 9},
		},
		{
			input:    []int{31, 41, 59, 26, 41, 58},
			k:        60,
			expected: []int{26, 31, 41, 41, 58, 59},
		},
	}
	for _, test := range tests {
		result := CountingSort(test.input, test.k)
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
