package main

import (
	"fmt"

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

	fmt.Printf("Is present: %t\n", list.Find(11))

	fmt.Printf("Size: %d\n", list.Size())

	list.Print()

	fmt.Printf("Removed: %t\n", list.Remove(0))
	fmt.Printf("Size: %d\n", list.Size())

	list.Print()

	list.AddFirst(7)

	fmt.Printf("Size: %d\n", list.Size())

	list.Remove(7)

	list.Print()

	fmt.Printf("Size: %d\n", list.Size())
}
