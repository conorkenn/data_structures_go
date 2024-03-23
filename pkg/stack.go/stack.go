package stack

import "errors"

type Stack struct {
	items []interface{}
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	top, _ := s.Top()
	s.items = s.items[:len(s.items)-1]
	return top, nil
}
