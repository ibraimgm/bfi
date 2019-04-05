package vm

import (
	"errors"
)

type stack struct {
	index int
	data  []int
}

func (s *stack) push(value int) {
	if s.index >= len(s.data) {
		s.data = append(s.data, value)
	} else {
		s.data[s.index] = value
	}

	s.index++
}

func (s *stack) peek() (int, error) {
	if s.index == 0 {
		return 0, errors.New("The stack is empty")
	}

	return s.data[s.index-1], nil
}

func (s *stack) pop() (int, error) {
	if s.index == 0 {
		return 0, errors.New("The stack is empty")
	}

	s.index--
	return s.data[s.index], nil
}

func newStack() *stack {
	return &stack{0, make([]int, 0)}
}
