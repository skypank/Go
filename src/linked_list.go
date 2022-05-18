package main

import "fmt"

// creating a user defined struct name LinkedListNode
type LinkedListNode struct {
	data     int
	nextNode *LinkedListNode
}

// contains two elements, a varialbe of type int
// and a pointer to next node

// this creates a new node and connect next node of list to this new node
func (ll *LinkedListNode) insertAtEnd(data int) {
	node := new(LinkedListNode)
	node.data = data
	ll.nextNode = node
}

// this creates a new node and connect this to existing start node and return new start node
func (ll *LinkedListNode) insertAtBeg(data int) (startNode *LinkedListNode) {
	node := new(LinkedListNode)
	node.data = data
	node.nextNode = ll
	return node
}

func (ll *LinkedListNode) deleteElement(data int) {
	node, prev := ll, ll
	for node != nil {
		if node.data == data {
			// case 1 : element to be deleted is at the head of the list
			if node == ll {
				// here the head of list has to be updated
				// we have to update the address of head list, since in go everything is pass by value
				// so explicit address update is needed
				*(ll) = *(node).nextNode
			} else {
				// case 2 : element to be deleted is at any other place
				// including last element
				prev.nextNode = node.nextNode
			}
		}
		prev = node
		node = node.nextNode
	}
}

func (ll *LinkedListNode) printList(msg string) {
	fmt.Println("The linked list:", msg)
	for ll.nextNode != nil {
		fmt.Printf("%d ->", ll.data)
		ll = ll.nextNode
	}
	fmt.Println(ll.data)
}

func main() {
	// Allocate memory for one one
	node := new(LinkedListNode)
	// add a new data element, the next node need not be set explicitly, as it will be set to nil by default
	node.data = 1
	// insert element at the end
	node.insertAtEnd(2)
	// insert element at the begining, the head of list will change, the added element will be new head
	node = node.insertAtBeg(3)
	// print the list
	node.printList("New list")
	// Delete a particular element from the list
	node.deleteElement(3)
	// print the list
	node.printList("After delete at the begining")
	node.deleteElement(2)
	// print the list
	node.printList("After delete at the end")
}
