package main

import "fmt"

// Declare Queue type as array of int
// type keyword allows you to define interfaces that can be either built in or user defined.
// Here we are using Queue as array of int
// Notice the capital letter of name, 'Q'ueue, this signifies that the interface can be exported to other packages
// If you dont want other packages to import your interface, use a small letter name of interface.
// This is how go allows you to make your interface private or public
type Queue []int

// This function can only be called via element of type Queue
func (q *Queue) enqueue(e int) {
	*q = append(*q, e)
}

func (q *Queue) deque() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	e := (*q)[0]
	*q = (*q)[1:]
	return e, true
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func main() {
	var Q Queue
	Q.enqueue(1)
	Q.enqueue(2)
	fmt.Println(Q.deque())
	fmt.Println(Q.deque())
	fmt.Println(Q.deque())
}
