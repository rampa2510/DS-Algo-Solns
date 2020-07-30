package main

import "fmt"

// Max2StackCapacity - The variable which is used to store the maximum capacity of 2 stacks
const Max2StackCapacity int = 20

// Stack2 - Implementation of 2 stacks in struct
// It consists of the arr and the top of stacks for both stack
// one stack will be top of stack (10-19) and the bottom stack (0-9)
type Stack2 struct {
	arr        [Max2StackCapacity]int
	tos1, tos2 int
}

// Push2 - This function is used to append elements to 2 stacks in one array
// it takes 2 arguments on the element which you want to insert
// and other a flag which indicates in which stack you want to append
// if flag = 1 then insert at the top stack else insert at bottom stack
func (s *Stack2) Push2(element, flag int) {
	if flag == 1 {

		if s.tos1 == 9 {
			fmt.Println("top stack is filled")
			return
		}

		s.tos1++
		s.arr[s.tos1] = element
		return
	}

	if s.tos2 == 19 {
		fmt.Println("bottom stack is filled")
		return
	}

	s.tos2++
	s.arr[s.tos2] = element
}

// IsEmpty2 - This method is used to check whether a stack is empty or not
// It accepts one argument flag which specifies for which array we eant to check whether it is empty or not
// if flag = 1 then check for the top stack else check for bottom stack
func (s *Stack2) IsEmpty2(flag int) bool {

	if flag == 1 {

		if s.tos1 == -1 {
			fmt.Println("Top Stack is empty")
			return true
		}

		return false
	}

	if s.tos2 == 9 {
		fmt.Println("Top Stack is empty")
		return true
	}
	return false

}

// Display2 - This method is used to display elements of one of the stack
// It takes one argument flag which specifies the stack of which you want to display the elements.
// if flag = 1 then display elements of the top stack else display elements of bottom stack
func (s *Stack2) Display2(flag int) {

	if s.IsEmpty2(flag) {
		return
	}

	if flag == 1 {

		for i := s.tos1; i >= 0; i-- {
			fmt.Println(s.arr[i])
		}
		return
	}

	for i := s.tos2; i >= 10; i-- {
		fmt.Println(s.arr[i])
	}
}

// Pop2 - This method is used to pop elements from the specified stack
// It takes one argument flag which indicate from which stack you want to pop elements
// if flag = 1 then pop elements of the top stack else pop elements of bottom stack
// It returns a integer pointer that points to the element poped
func (s *Stack2) Pop2(flag int) *int {
	if s.IsEmpty2(flag) {
		return nil
	}

	if flag == 1 {
		i := s.tos1
		s.tos1--
		return &s.arr[i]
	}
	i := s.tos2
	s.tos2--
	return &s.arr[i]
}
