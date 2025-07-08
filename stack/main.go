// Stacks
// A stack is a linear data structure that follows a particular order in which the operations are performed. The order may be LIFO (Last In First Out) or FILO (First In Last Out). Mainly three basic operations are performed in the stack:

// Push: adds an element to the collection.

// Pop: removes an element from the collection. A pop can result in stack underflow if the stack is empty.

// Peek or Top: returns the top item without removing it from the stack.

// The basic principle of stack operation is that in a stack, the element that is added last is the first one to come off, thus the name "Last in First Out".

package main

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func (s *Stack) Push(data int) {
	s.Items = append(s.Items, data)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	topItem := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]

	return topItem
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	topItem := s.Items[len(s.Items)-1]

	return topItem
}

func main() {
	s := Stack{}

	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	fmt.Println(s)
}
