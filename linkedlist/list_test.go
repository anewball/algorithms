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

// This test inserted a value at the end of an empty linkedlist
func TestAddLastToEmptyList(t *testing.T) {
	list := New()

	const x = 5

	list.AddLast(x)

	if !list.Find(x) {
		t.Errorf("The cardinal value %d was not been iserted", x)
	}
}

// This test add a new element at the end of a non empty linkedlist
func TestAddLastToNonEmptyList(t *testing.T) {
	list := getList()

	const x = 11

	list.AddLast(x)

	if !list.Find(x) {
		t.Errorf("The cardinal value %d was not been inserted into the linkedlist", x)
	}
}

// This test use the AddBefore method to add an element
// into a linkedlist
func TestAddElementAtBeginin(t *testing.T) {
	list := getList()

	const x = 0
	const value = 1

	if !list.AddBefore(value, x) {
		t.Errorf("The cardinal value %d is not in the linkedlist", value)
	}
}

// This test and an element in between another element
func TestAddElementBefore(t *testing.T) {
	list := getList()

	const x = 8
	const target = 7

	if !list.AddBefore(target, x) {
		t.Errorf("The cardinal value %d is not in the linkedlist", target)
	}
}

// This test try to add a value before a non existing one
func TestAddBeforeNonExistingElement(t *testing.T) {
	list := getList()

	const x = 10
	const target = 9

	if list.AddBefore(target, x) {
		t.Errorf("The cardinal value %d is in the linkedlist", target)
	}
}

// This test try to add a non existing element into an empty linkedlist
func TestAddBeforeAnEmptyList(t *testing.T) {
	list := New()

	const x = 10
	const target = 9

	if list.AddBefore(target, x) {
		t.Errorf("The cardinal value %d is in the linkedlist", target)
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
