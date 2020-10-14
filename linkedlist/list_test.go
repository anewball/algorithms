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

// This test validated that the number of elements inside the
// the linkedlist equals 7. Otherwise it will return an error.
func TestListSizeEqualSeven(t *testing.T) {
	list := getList()

	const n = 7

	if list.Size() != n {
		t.Errorf("The size of the linkedlist is %d.", list.Size())
	}
}

// This test removed an element from a linkedlist. If the element
// was removed, true will be returned otherwise false.
func TestRemoveAnElementFromList(t *testing.T) {
	list := getList()

	const element = 4

	if !list.Remove(element) {
		t.Errorf("The element %d is not part of the list.", element)
	}
}

// This test try to remove an element that is not part of the linkedlist.
func TestRemoveNonExistingElementFromList(t *testing.T) {
	list := getList()

	const element = 100

	if list.Remove(element) {
		t.Errorf("The element %d is part of the list.", element)
	}
}

// This test removed the last element from a linkedlist.
// Thereby, the current and the tail should be pointing to the same node.
func TestRemoveLastElementFromList(t *testing.T) {
	list := getList()

	const x = 7

	if !list.Remove(x) {
		t.Errorf("The cardinal value %d is not part of the linkedlist.", x)
	}
}

// This test removed the only element inside the linkedlist.
func TestRemoveOnlyElementFromList(t *testing.T) {
	list := New()

	const x = 1

	list.Add(x)

	if !list.Remove(x) {
		t.Errorf("The cardinal value %d is not part of the linkedlist.", x)
	}
}

// This test try to remove an element from a empty linkedlist.
func TestTryToRemoveElementFromEmptyList(t *testing.T) {
	list := New()

	const x = 1

	if list.Remove(x) {
		t.Errorf("The cardinal value %d is part of the list.", x)
	}
}

// This test printed all the elements inside a linkedlist
func TestPrintList(t *testing.T) {
	list := getList()

	list.Print()
}

// This test searched for an element inside the linkedlist.
func TestFindElementInList(t *testing.T) {
	list := getList()

	const x = 5

	if !list.Find(x) {
		t.Errorf("The value %d is not part of the linkedlist.", x)
	}
}

// This test try to find an element inside an empty linkedlist.
func TestFindElementInEmptyList(t *testing.T) {
	list := New()

	const x = 3

	if list.Find(x) {
		t.Errorf("The cardinal value %d is part of the linkedlist.", x)
	}
}

// This test added an element at the begin of an existing linkedlist.
func TestAddElementAtBeginOfExistingList(t *testing.T) {
	list := getList()

	const x = 8

	list.AddFirst(x)
}

// This test added a value to an empty linkedlist.
func TestAddElementAtBeginOfANewList(t *testing.T) {
	list := New()

	const x = 1

	list.AddFirst(x)
}
