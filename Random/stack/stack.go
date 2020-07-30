package main

import "fmt"

// MaxStackCapacity - The variable which is used to store the maximum capacity of stack
const MaxStackCapacity int = 10

// Stack - Struct implementation of stack arr
type Stack struct {
	arr [MaxStackCapacity]int
	tos int
}

// Push - Push method for stack
// It takes one argument and that is the element you want to insert
func (s *Stack) Push(element int) {

	if s.tos+1 == MaxStackCapacity {
		fmt.Println("Stack capacity reached")
		return
	}
	s.tos++
	s.arr[s.tos] = element
}

// IsEmpty - Check if Stack empty or not
func (s *Stack) IsEmpty() bool {
	return s.tos == -1
}

// Display - Display all the elements of Stack
func (s *Stack) Display() {

	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	for i := s.tos; i >= 0; i-- {
		fmt.Println(s.arr[i])
	}
}

// Pop - Pop method for Stack
// It returns a integer pointer which points to the element that is popped
func (s *Stack) Pop() *int {
	if s.IsEmpty() {
		fmt.Println("Empty stack")
		return nil
	}
	idx := s.tos
	s.tos--
	return &s.arr[idx]
}
