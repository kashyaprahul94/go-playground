package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kashyaprahul94/go-playground/pkg/algorithm"
	datamodel "github.com/kashyaprahul94/go-playground/pkg/data-model"
	_ "github.com/kashyaprahul94/go-playground/pkg/web/mux"
	_ "github.com/kashyaprahul94/go-playground/pkg/web/native"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func algos() {

	// Sorting

	sorter := algorithm.BubbleSort{Numbers: []int{4, 3, 1, 5, 2}}
	sortingResult := sorter.Sort()

	fmt.Printf("%v", sortingResult)
	fmt.Println()

	// LinkedList

	list := algorithm.SinglyLinkedList()

	list.Push(4)
	list.Push(3)
	list.Push(1)

	list.Push(12.2)
	list.Push("s")

	fmt.Println(list)
}

func webStuff() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4444"
	}

	// native.StartServer(port)
	// mux.StartServer(port)
}

func dataModels() {

	datamodel.PlayWithJSON()
}

func run() error {

	// algos()
	// webStuff()
	// dataModels()

	return nil
}
