package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type Node[T any] struct {
	Data T
	Next *Node[T]
}

func (n *Node[T]) GetData() T {
	return n.Data
}

type SinglyLinkedList[T any] struct {
	Head *Node[T]
}

func (ll *SinglyLinkedList[T]) InsertAtHead(data T) {
	newNode := &Node[T]{Data: data, Next: ll.Head}
	ll.Head = newNode
}

func (ll *SinglyLinkedList[T]) InsertAfter(target, data T) {
	current := ll.Head

	for current != nil && !cmp.Equal(current.Data, target) {
		current = current.Next
	}

	if current != nil {
		newNode := &Node[T]{Data: data, Next: current.Next}
		current.Next = newNode
	}
}

func (ll *SinglyLinkedList[T]) InsertAtTail(data T) {
	newNode := &Node[T]{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *SinglyLinkedList[T]) Display() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *SinglyLinkedList[T]) Search(target T) *Node[T] {
	current := ll.Head
	for current != nil {
		if cmp.Equal(current.Data, target) {
			return current
		}

		current = current.Next
	}
	return nil
}

func (ll *SinglyLinkedList[T]) DeleteNode(target T) {
	if ll.Head == nil {
		return
	}

	if cmp.Equal(ll.Head.Data, target) {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil && !cmp.Equal(current.Next.Data, target) {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func main() {
	ll := SinglyLinkedList[int]{}

	ll.InsertAtHead(10)
	ll.InsertAtHead(20)

	ll.InsertAfter(20, 25)

	ll.InsertAtHead(30)
	ll.InsertAtTail(40)
	ll.InsertAtTail(50)

	ll.InsertAfter(40, 30)

	fmt.Println("Linked List:")
	ll.Display() // Output: 30 -> 20 -> 10 -> 40 -> 50 -> nil

	searchNode := ll.Search(10)
	if searchNode != nil {
		fmt.Println("Node ditemukan:", searchNode.GetData()) // Output: Node ditemukan: 10
	} else {
		fmt.Println("Node tidak ditemukan")
	}

	ll.DeleteNode(20)
	fmt.Println("Setelah menghapus 20:")
	ll.Display()
}
