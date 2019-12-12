package main

import (
	"log"
	"os"

	"github.com/kashyaprahul94/go-playground/pkg/algorithm"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {

	list := algorithm.SinglyLinkedList()

	list.Push(4)
	list.Push(3)
	list.Push(1)

	list.Push(12.2)
	list.Push("s")

	list.Print()

	return nil
}
