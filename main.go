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

	list.AddBefore(6, 7)
	list.AddBefore(12, 1.3)

	list.Print()
}
