package selection

import "testing"

func TestBucketSort(t *testing.T) {
	tests := []struct {
		input       []int
		expectedMin int
		expectedMax int
	}{
		{
			input:       []int{1, 2, 3},
			expectedMin: 1,
			expectedMax: 3,
		},
		{
			input:       []int{3, 2, 1},
			expectedMin: 1,
			expectedMax: 3,
		},
		{
			input:       []int{1, 2, 3, 7, 6, 9, 4},
			expectedMin: 1,
			expectedMax: 9,
		},
		{
			input:       []int{31, 41, 59, 26, 41, 58},
			expectedMin: 26,
			expectedMax: 59,
		},
	}
	for _, test := range tests {
		max, min := MinMaxSelection(test.input)
		if max != test.expectedMax {
			t.Fatalf("Test Failed: expected %v, instead recieved: %v",
				test.expectedMax, max)
		}
		if min != test.expectedMin {
			t.Fatalf("Test Failed: expected %v, instead recieved: %v",
				test.expectedMin, min)
		}
	}
}
