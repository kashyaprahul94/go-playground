package algorithm

import "fmt"

// Item is the kind of item to be stored
type Item interface{}

// node is awesome too
type node struct {
	value Item
	next  *node
}

// list is awesome too
type list struct {
	head *node
}

func push(l *list, nodeToInsert *node) {

	if l.head == nil {
		l.head = nodeToInsert
	} else {

		currentNode := l.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = nodeToInsert
	}
}

// Push is used to insert element at the end of the list
func (l *list) Push(element Item) {
	newNode := &node{value: element, next: nil}

	push(l, newNode)
}

// Print is used to print the contents of list
func (l *list) Print() {
	head := l.head

	fmt.Printf("\nHead")

	for head != nil {
		fmt.Printf(" -> %v", head.value)
		head = head.next
	}

	fmt.Print(" -> nil\n\n")
}

// SinglyLinkedList is awesome
func SinglyLinkedList() *list {

	newList := new(list)

	return &(*(newList))
}
