package linkedlist

import "fmt"

type node struct {
	data interface{}
	next *node
}

// List list of node
type List struct {
	head *node
	tail *node
	size int
}

// New create a new LinkedList
func New() *List {
	return new(List)
}

// Add add a new node to the linkedlist
func (l *List) Add(data interface{}) {
	if l.size == 0 {
		l.head = &node{data: data}
		l.tail = l.head
	} else {
		current := l.head
		for {
			if current.next == nil {
				current.next = &node{data: data}
				l.tail = current.next
				break
			}
			current = current.next
		}
	}
	l.size++
}

// AddFirst add a new node at the beginning at the linkedlist
func (l *List) AddFirst(data interface{}) {
	if l.size > 0 {
		temp := l.head
		l.head = &node{data: data}
		l.head.next = temp
	} else {
		l.head = &node{data: data}
		l.tail = l.head
	}

	l.size++
}

// Find find search into the linkedlist for a
// specific value and return true if value exist otherwise false
func (l *List) Find(data interface{}) bool {
	current := l.head
	_, err := find(current, data)
	return err == nil
}

// Size return number of elements stored inside the linkedlist
func (l *List) Size() int {
	return l.size
}

// Print the content of the linkedlist
func (l *List) Print() {
	current := l.head
	print(current)
}

// Remove remove an element from the linkedlist
func (l *List) Remove(data interface{}) bool {
	previous := l.head
	current := l.head

	if l.size > 0 {
		for current.data != data {
			if current.next == nil {
				return false
			}
			previous = current
			current = current.next
		}

		if current == l.head {
			l.head = current.next
		} else {
			previous.next = current.next
			if current == l.tail {
				l.tail = previous
			}
		}

		l.size--

		return true
	}

	return false
}

func find(root *node, data interface{}) (*node, error) {
	if root != nil {
		if root.data == data {
			return root, nil
		}
		return find(root.next, data)
	}
	return &node{}, fmt.Errorf("the %v does not exist", data)
}

func print(root *node) {
	if root != nil {
		fmt.Println(root.data)
		print(root.next)
	}
}
