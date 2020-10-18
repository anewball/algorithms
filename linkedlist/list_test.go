package linkedlist

import (
	"testing"
)

func getList() *List {
	list := New()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)

	return list
}

func TestSize(t *testing.T) {
	// This test validated that the number of elements inside the
	// the linkedlist equals 7. Otherwise it will return an error.
	list := getList()

	newNode := 7

	if list.Size() != newNode {
		t.Errorf("The size of the linkedlist is %d.", list.Size())
	}
}

func TestRemove(t *testing.T) {
	// This test try to remove an element from a empty linkedlist.
	list := New()

	newNode := 1

	if list.Remove(newNode) {
		t.Errorf("The cardinal value %d is part of the list.", newNode)
	}

	// This test removed the only element inside the linkedlist.
	list = New()

	newNode = 1

	list.Add(newNode)
	if !list.Remove(newNode) {
		t.Errorf("The cardinal value %d is not part of the linkedlist.", newNode)
	}

	// This test removed the last element from a linkedlist.
	// Thereby, the current and the tail should be pointing to the same node.
	list = getList()

	newNode = 7

	if !list.Remove(newNode) {
		t.Errorf("The cardinal value %d is not part of the linkedlist.", newNode)
	}

	// This test try to remove an element that is not part of the linkedlist.
	list = getList()

	newNode = 100

	if list.Remove(newNode) {
		t.Errorf("The element %d is part of the list.", newNode)
	}

	// This test removed an element from a linkedlist. If the element
	// was removed, true will be returned otherwise false.
	list = getList()

	newNode = 4

	if !list.Remove(newNode) {
		t.Errorf("The element %d is not part of the list.", newNode)
	}
}

// This test printed all the elements inside a linkedlist
func TestPrint(t *testing.T) {
	list := getList()

	list.Print()
}

func TestFind(t *testing.T) {
	// This test try to find an element inside an empty linkedlist.
	list := New()

	newNode := 3

	if list.Find(newNode) {
		t.Errorf("The cardinal value %d is part of the linkedlist.", newNode)
	}

	// This test searched for an element inside the linkedlist.
	list = getList()

	newNode = 5

	if !list.Find(newNode) {
		t.Errorf("The value %d is not part of the linkedlist.", newNode)
	}
}

func TestAddLast(t *testing.T) {
	// This test add a new element at the end of a non empty linkedlist
	list := getList()

	newNode := 11

	list.AddLast(newNode)
	if !list.Find(newNode) {
		t.Errorf("The cardinal value %d was not been inserted into the linkedlist", newNode)
	}

	// This test inserted a value at the end of an empty linkedlist
	list = New()

	newNode = 5

	list.AddLast(newNode)
	if !list.Find(newNode) {
		t.Errorf("The cardinal value %d was not been iserted", newNode)
	}

	// This test added a value to an empty linkedlist.
	list = New()

	newNode = 1

	list.AddFirst(newNode)
	if !list.Find(newNode) {
		t.Errorf("The cardinal value %d was not been iserted", newNode)
	}

	// This test added an element at the begin of an existing linkedlist.
	list = getList()

	newNode = 8

	list.AddFirst(newNode)
	if !list.Find(newNode) {
		t.Errorf("The cardinal value %d was not been iserted", newNode)
	}
}

func TestAddBefore(t *testing.T) {
	// This test try to add a non existing element into an empty linkedlist
	list := New()

	newNode := 10
	target := 9

	if list.AddBefore(target, newNode) {
		t.Errorf("The cardinal value %d is in the linkedlist", target)
	}

	// This test try to add a value before a non existing one
	list = getList()

	newNode = 10
	target = 9

	if list.AddBefore(target, newNode) {
		t.Errorf("The cardinal value %d is in the linkedlist", target)
	}

	// This test and an element in between another element
	list = getList()

	newNode = 8
	target = 7

	if !list.AddBefore(target, newNode) {
		t.Errorf("The cardinal value %d is not in the linkedlist", target)
	}

	// This test use the AddBefore method to add an element
	// into a linkedlist
	list = getList()

	newNode = 0
	target = 1

	if !list.AddBefore(target, newNode) {
		t.Errorf("The cardinal value %d is not in the linkedlist", target)
	}
}

func TestAddAfter(t *testing.T) {
	list := getList()

	// Testing a non empty list with an existing element.
	// The AddAfter method should return true

	newNode := 1
	targetNode := 3

	if !list.AddAfter(targetNode, newNode) {
		t.Errorf("The cardinal value %d is not in the linkedlist", targetNode)
	}

	// Testing a non empty list with an existing element. This new node will
	// be inserted after the existing one, at the end of the linkedlist. It
	// will become the tail

	newNode = 1
	targetNode = 7

	if !list.AddAfter(targetNode, newNode) {
		t.Errorf("The cardinal value %d is not in the linkedlist", targetNode)
	}

	// Testing a non empty list with a non existing element.
	// AddAfter should return false since the targetNode is not part of the set
	// of elements

	newNode = 1
	targetNode = 8

	if list.AddAfter(targetNode, newNode) {
		t.Errorf("The cardinal value %d is in the linkedlist", targetNode)
	}

	// Testing an empty list. Try to insert a new node after one that
	// is not part of the empty list.

	newNode = 1
	targetNode = 8

	list = New()

	if list.AddAfter(targetNode, newNode) {
		t.Errorf("The cardinal value %d is in the linkedlist", targetNode)
	}
}
