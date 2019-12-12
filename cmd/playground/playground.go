package main

import "github.com/kashyaprahul94/go-playground/pkg/algorithm"

func main() {

	list := algorithm.SinglyLinkedList()

	list.Push(4)
	list.Push(3)
	list.Push(1)

	list.Push(12.2)
	list.Push("s")

	list.Print()
}
