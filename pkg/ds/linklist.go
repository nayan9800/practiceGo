package ds

import (
	"fmt"
)

/*Linked list struct stores Head node*/
type LinkedList struct {
	Head *node
}

/*node stores value and pointer to next node*/
type node struct {
	Value int
	next  *node
}

/*Creates new linklist*/
func NewLinkedList() LinkedList { return LinkedList{Head: nil} }

/*Inserts new element in Linked list*/
func (ll *LinkedList) Insert(value int) {
	nn := node{Value: value, next: nil}
	if ll.Head == nil {
		ll.Head = &nn
	} else {

		current := ll.Head
		for current.next != nil {
			current = current.next
		}
		current.next = &nn
	}
}

/*Traverse through linked list*/
func (ll *LinkedList) Traverse() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.next
	}
}

/*search element in linked list*/
func (ll *LinkedList) Search(val int) (ans bool) {
	current := ll.Head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.next
	}
	return
}

/*Delete element in linked list*/
func (ll *LinkedList) Delete(val int) (ans bool) {
	current := ll.Head
	for current != nil {
		if current.next.Value == val {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return
}
