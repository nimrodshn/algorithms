package linkedlist

import (
	"testing"
)

var (
	l = New()
)

func TestLinkedList(t *testing.T) {
	tests := []struct {
		testCase    func() error
		expectedErr error
	}{
		{
			testCase: func() error {
				l.Add(10)
				_, err := l.Search(10)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			testCase: func() error {
				l.Add(22)
				_, err := l.Search(22)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			testCase: func() error {
				err := l.Remove(10)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			testCase: func() error {
				// Expected ErrNotFound.
				err := l.Remove(4)
				if err != nil {
					return err
				}
				return nil
			},
			expectedErr: ErrNotFound,
		},
		{
			testCase: func() error {
				// Expected ErrNotFound.
				_, err := l.Search(600)
				if err != nil {
					return err
				}
				return nil
			},
			expectedErr: ErrNotFound,
		},
	}
	for _, test := range tests {
		err := test.testCase()
		if err != test.expectedErr {
			t.Fatalf("Expected an error: %v, instead recieved %v", test.expectedErr.Error(), err.Error())
		}
	}
}
