package vm

import (
	"testing"
)

func TestPushPop(t *testing.T) {
	testCases := []struct {
		value       int
		shouldPush  bool
		shouldError bool
	}{
		{value: 0, shouldPush: false, shouldError: true},
		{value: 3, shouldPush: true},
		{value: 9, shouldPush: true},
		{value: 7, shouldPush: true},
		{value: 7, shouldPush: false},
		{value: 15, shouldPush: true},
		{value: 15, shouldPush: false},
		{value: 9, shouldPush: false},
		{value: 3, shouldPush: false},
		{value: 0, shouldPush: false, shouldError: true},
	}

	s := newStack()

	for i, test := range testCases {
		if test.shouldPush {
			s.push(test.value)
			peek, err := s.peek()

			if err != nil {
				t.Errorf("Case %v, unexpected error received: \"%v\"", i, err)
			}

			if peek != test.value {
				t.Errorf("Case %v, expected peek value to be \"%v\", received \"%v\"", i, test.value, peek)
			}

		} else {
			v, err := s.pop()

			if v != test.value {
				t.Errorf("Case %v, expected pop value \"%v\", received \"%v\"", i, test.value, v)
			}

			if test.shouldError && err == nil {
				t.Errorf("Case %v, expected pop to raise error, received nil", i)
			}

			if !test.shouldError && err != nil {
				t.Errorf("Case %v, unexpected error received: \"%v\"", i, err)
			}
		}
	}
}

func TestEmptyPeek(t *testing.T) {
	s := newStack()

	if _, err := s.peek(); err == nil {
		t.Errorf("Expected error when peeking from empty stack")
	}

	s.push(1)
	s.push(2)

	if _, err := s.peek(); err != nil {
		t.Errorf("Unexpected error when peeking stack with values: \"%v\"", err)
	}

	s.pop()
	if _, err := s.peek(); err != nil {
		t.Errorf("Unexpected error when peeking stack with values: \"%v\"", err)
	}

	s.pop()
	if _, err := s.peek(); err == nil {
		t.Errorf("Expected error when peeking from empty stack")
	}
}
