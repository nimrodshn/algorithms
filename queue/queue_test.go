package queue

import (
	"testing"
)

var (
	length = 3
	s      = New(length)
)

func TestQueue(t *testing.T) {
	tests := []struct {
		testCase    func() error
		expectedErr error
	}{
		{
			testCase: func() error {
				err := s.Insert(10)
				if err != nil {
					return err
				}
				if s.Size() != 1 {
					t.Fatalf("Expected queue size to be %d instead recieved %d", 1, s.Size())
				}
				return nil
			},
		},
		{
			testCase: func() error {
				err := s.Insert(10)
				if err != nil {
					return err
				}
				if s.Size() != 2 {
					t.Fatalf("Expected queue size to be %d instead recieved %d", 2, s.Size())
				}
				return nil
			},
		},
		{
			testCase: func() error {
				x, err := s.Remove()
				if err != nil {
					return err
				}
				if s.Size() != 1 {
					t.Fatalf("Expected queue size to be %d instead recieved %d", 1, s.Size())
				}
				if x != 10 {
					t.Fatalf("Expected removed element to be %d instead recieved %d", 10, x)
				}
				return nil
			},
		},
		{
			testCase: func() error {
				x, err := s.Remove()
				if err != nil {
					return err
				}
				if s.Size() != 0 {
					t.Fatalf("Expected queue size to be %d instead recieved %d", 0, s.Size())
				}
				if x != 10 {
					t.Fatalf("Expected removed element to be %d instead recieved %d", 10, x)
				}
				return nil
			},
		},
		{
			testCase: func() error {
				_, err := s.Remove()
				if err != nil {
					return err
				}
				return nil
			},
			expectedErr: ErrQueueUnderflow,
		},
		{
			testCase: func() error {
				for i := 0; i <= length; i++ {
					err := s.Insert(i)
					if err != nil {
						return err
					}
				}
				return nil
			},
			expectedErr: ErrQueueOverflow,
		},
	}
	for _, test := range tests {
		err := test.testCase()
		if err != test.expectedErr {
			t.Fatalf("Expected an error: %v, instead recieved %v", test.expectedErr.Error(), err.Error())
		}
	}
}
