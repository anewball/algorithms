package main

import (
	"github.com/algorithms/linkedlist"
)

func main() {
	list := linkedlist.New()

	list.AddFirst(2)
	list.AddFirst(1)

	list.Print()
}
