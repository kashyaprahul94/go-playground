package algorithm

import "fmt"

// Item is the kind of item to be stored
type Item interface{}

// linkedNode is awesome too
type linkedNode struct {
	value Item
	next  *linkedNode
}

// LinkedList is awesome too
type LinkedList struct {
	head *linkedNode
}

func push(l *LinkedList, node *linkedNode) {

	if l.head == nil {
		l.head = node
	} else {

		currentNode := l.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = node
	}
}

// Push is used to insert element at the end of the list
func (l *LinkedList) Push(element Item) {
	node := &linkedNode{value: element, next: nil}

	push(l, node)
}

// Print is used to print the contents of list
func (l *LinkedList) String() string {

	result := ""

	head := l.head

	result += fmt.Sprintf("\nHead")

	for head != nil {
		result += fmt.Sprintf(" -> %v", head.value)
		head = head.next
	}

	result += fmt.Sprintf(" -> nil\n")

	return result
}

// SinglyLinkedList is awesome
func SinglyLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}
