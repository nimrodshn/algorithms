package linkedlist

import (
	"fmt"
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
		{
			testCase: func() error {
				err := l.Remove(22)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			testCase: func() error {
				for i := 0; i < 10; i++ {
					l.Add(i)
				}
				// The linked lis should now look like: head->9,8,7...,0
				i := 9
				err := l.Iterate(func(l *Link) error {
					if l.key != i {
						return fmt.Errorf("Expected %d, instead recieved %d", l.key, i)
					}
					i--
					return nil
				})
				if err != nil {
					return err
				}
				// Reverse the linked list.
				l.Reverse()
				// Now it should look like: head->0,1,...,9
				i = 0
				err = l.Iterate(func(l *Link) error {
					if l.key != i {
						return fmt.Errorf("Expected %d, instead recieved %d", l.key, i)
					}
					i++
					return nil
				})
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	for _, test := range tests {
		err := test.testCase()
		if err != test.expectedErr {
			t.Fatalf("Expected an error: %v, instead recieved %v", test.expectedErr.Error(), err.Error())
		}
	}
}
