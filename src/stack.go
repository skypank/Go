package main

import (
	"fmt"
)

// Declare stack type as array of int
// type keyword allows you to define interfaces that can be either built in or user defined.
// Here we are using Stack as array of int
// Notice the capital letter of name, 'S'tack, this signifies that the interface can be exported to other packages
// If you dont want other packages to import your interface, use a small letter name of interface.
// This is how go allows you to make your interface private or public

type Stack []int

// This function can be called only via element of type Stack
// This is how go enforces method encapsulation, since this function cannot be called by any other interface
func (s *Stack) push(e int) {
	// add the element on stack, which is array of int
	*s = append(*s, e)
}

func (s *Stack) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e, true
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) top() int {
	if !s.isEmpty() {
		return (*s)[len(*s)-1]
	}
	return -1
}

func main() {
	// declaring a varialbe of type Stack
	var s Stack

	// the underlying functions can be called via Stack type varilable only
	s.push(1)
	s.push(2)
	fmt.Println(s.pop())
	s.push(3)
	fmt.Println(s.top())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.top())
}
