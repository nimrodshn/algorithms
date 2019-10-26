package stack

import (
	"errors"
)

var (
	// ErrStackUnderflow is an error indicating the stack is empty.
	ErrStackUnderflow = errors.New("Stack Underflow")

	// ErrStackOverflow is an error indicating the stack is full.
	ErrStackOverflow = errors.New("Stack Overflow")
)

// Stack is an implementation of a stack.
type Stack struct {
	length int
	top    int
	array  []int
}

// New initializes a stack of size n.
func New(n int) *Stack {
	return &Stack{
		length: n,
		top:    -1,
		array:  make([]int, n),
	}
}

// Pop pops the top element off the stack
func (s *Stack) Pop() (int, error) {
	if s.empty() {
		return -1, ErrStackUnderflow
	}
	result := s.array[s.top]
	s.top--
	return result, nil
}

// Push pushes a new element to the array.
func (s *Stack) Push(x int) error {
	if s.full() {
		return ErrStackOverflow
	}
	s.top++
	s.array[s.top] = x
	return nil
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) empty() bool {
	return s.top == -1
}

func (s *Stack) full() bool {
	return s.top == (s.length - 1)
}
