package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kashyaprahul94/go-playground/pkg/algorithm"
	"github.com/kashyaprahul94/go-playground/pkg/web"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func algos() {
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

	web.StartServer(port)
}

func run() error {

	// algos()
	webStuff()

	return nil
}
