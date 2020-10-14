package main

import (
	"github.com/algorithms/linkedlist"
)

func main() {
	list := linkedlist.New()

	list.Add(6)
	list.Add(5)
	list.Add(4)
	list.Add(3)
	list.Add(2)
	list.Add(1)
	list.Add(0)

	list.Print()
}
