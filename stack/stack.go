package stack

import "errors"

// Stack is a simple LIFO stack for any type.
type Stack []interface{}

// Push adds an item to the top of the stack
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

// Pop removes the top item from the stack and returns it. An error is returned
// if called on an empty stack.
func (s *Stack) Pop() (interface{}, error) {
	length := s.Length()
	
	if length == 0 {
		return nil, errors.New("cannot pop from an empty stack")
	}
	
	item := (*s)[length-1]
	*s = (*s)[:length-1]
	
	return item, nil
}

// Length is a wrapper for `len(s)`
func (s *Stack) Length() int {
	return len(*s)
}

// IsEmpty returns true if the stack is empty, and false otherwise
func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}