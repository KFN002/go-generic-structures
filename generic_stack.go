package main

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *Stack[T]) Pop() T {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}
