package main

import (
	"github.com/algorithms/linkedlist"
)

func main() {
	list := linkedlist.New()

	list.Add(1)
	list.Add(3)

	list.AddBefore(3, 2)

	list.Print()
}
