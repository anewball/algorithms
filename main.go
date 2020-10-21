package main

import (
	"fmt"

	"github.com/algorithms/linkedlist"
)

func main() {
	list := linkedlist.New()

	list.AddFirst(2)
	list.AddFirst(1)

	list.Print()

	fmt.Println()

	list.RemoveLast()

	list.Print()

	fmt.Println()

	list.RemoveLast()
}
