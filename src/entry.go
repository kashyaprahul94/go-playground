package main

import (
	"algorithm"
)

func main() {

	// numbers := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}

	// sorter := algorithm.InsertionSort{Numbers: numbers}

	// var sortResult = sorter.Sort()

	// fmt.Println("")
	// fmt.Println(sortResult)
	// fmt.Println("")

	list := algorithm.SinglyLinkedList()

	list.Push(4)
	list.Push(3)
	list.Push(1)

	list.Push(12.2)
	list.Push("s")

	list.Print()
}
