package stack

import (
	"testing"
)

var (
	length = 3
	s      = New(length)
)

func TestStack(t *testing.T) {
	tests := []struct {
		testCase    func() error
		expectedErr error
	}{
		{
			testCase: func() error {
				err := s.Push(10)
				if err != nil {
					return err
				}
				if s.Size() != 1 {
					t.Fatalf("Expected stack sized to be %d instead recieved %d", 1, s.Size())
				}
				return nil
			},
		},
		{
			testCase: func() error {
				err := s.Push(10)
				if err != nil {
					return err
				}
				if s.Size() != 2 {
					t.Fatalf("Expected stack sized to be %d instead recieved %d", 2, s.Size())
				}
				return nil
			},
		},
		{
			testCase: func() error {
				x, err := s.Pop()
				if err != nil {
					return err
				}
				if s.Size() != 1 {
					t.Fatalf("Expected stack sized to be %d instead recieved %d", 1, s.Size())
				}
				if x != 10 {
					t.Fatalf("Expected popped element to be %d instead recieved %d", 10, s.Size())
				}
				return nil
			},
		},
		{
			testCase: func() error {
				x, err := s.Pop()
				if err != nil {
					return err
				}
				if s.Size() != 0 {
					t.Fatalf("Expected stack sized to be %d instead recieved %d", 0, s.Size())
				}
				if x != 10 {
					t.Fatalf("Expected popped element to be %d instead recieved %d", 10, s.Size())
				}
				return nil
			},
		},
		{
			testCase: func() error {
				_, err := s.Pop()
				if err != nil {
					return err
				}
				return nil
			},
			expectedErr: ErrStackUnderflow,
		},
		{
			testCase: func() error {
				for i := 0; i <= length; i++ {
					err := s.Push(i)
					if err != nil {
						return err
					}
				}
				return nil
			},
			expectedErr: ErrStackOverflow,
		},
	}
	for _, test := range tests {
		err := test.testCase()
		if err != test.expectedErr {
			t.Fatalf("Expected an error: %v, instead recieved %v", test.expectedErr.Error(), err.Error())
		}
	}
}
