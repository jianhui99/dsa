package main

import "fmt"

type Node struct {
	data int
	Prev *Node
	Next *Node
}

type LinkedList struct {
	Head    *Node
	Current *Node
}

func NewLinkedList(data int) *LinkedList {
	initNode := &Node{
		data: data,
	}

	return &LinkedList{
		Head:    initNode,
		Current: initNode,
	}
}

func (ll *LinkedList) InsertAfter(data int) {
	newNode := &Node{
		data: data,
		Prev: ll.Current,
	}

	newNode.Next = ll.Current.Next

	if ll.Current.Next != nil {
		ll.Current.Next.Prev = newNode
	}

	ll.Current.Next = newNode
	ll.Current = newNode
}

func (ll *LinkedList) PrintForward() {
	for node := ll.Head; node != nil; node = node.Next {
		fmt.Println(node.data)
	}
}
func (ll *LinkedList) PrintBackward() {
	for node := ll.Current; node != nil; node = node.Prev {
		fmt.Println(node.data)
	}
}

func main() {
	ll := NewLinkedList(1)

	ll.InsertAfter(2)
	ll.InsertAfter(3)
	ll.InsertAfter(4)
	ll.InsertAfter(5)

	ll.PrintForward()
	ll.PrintBackward()
}
