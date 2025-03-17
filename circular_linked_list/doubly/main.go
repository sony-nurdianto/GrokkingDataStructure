package main

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

type CircularDoublyLinkedList[T cmp.Ordered] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (dl *CircularDoublyLinkedList[T]) InsertAtHead(data T) {
	newNode := &Node[T]{Data: data}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
		dl.Head.Next = dl.Head
		dl.Head.Prev = dl.Tail
		dl.Tail.Next = dl.Head
		return
	}

	newNode.Next = dl.Head
	newNode.Prev = dl.Tail
	dl.Head.Prev = newNode
	dl.Tail.Next = newNode
	dl.Head = newNode
}

func (dl *CircularDoublyLinkedList[T]) InsertAtTail(data T) {
	newNode := &Node[T]{Data: data}

	if dl.Tail == nil {
		dl.InsertAtHead(data)
		return
	}

	newNode.Prev = dl.Tail
	newNode.Next = dl.Head
	dl.Tail.Next = newNode
	dl.Head.Prev = newNode
	dl.Tail = newNode
}

func (dl *CircularDoublyLinkedList[T]) DisplayFromHead() {
	if dl.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := dl.Head
	fmt.Print("Forward: ")
	for {
		fmt.Printf("[%v] ⇄ ", current.Data)
		current = current.Next
		if current == dl.Head {
			break
		}
	}
	fmt.Println(" (circular back to Head)")
}

func (dl *CircularDoublyLinkedList[T]) DisplayFromTail() {
	if dl.Tail == nil {
		fmt.Println("List is empty")
		return
	}

	current := dl.Tail
	fmt.Print("Backward: ")
	for {
		fmt.Printf("[%v] ⇄ ", current.Data)
		current = current.Prev
		if current == dl.Tail {
			break
		}
	}
	fmt.Println(" (circular back to Tail)")
}

func main() {
	dl := CircularDoublyLinkedList[int]{}

	dl.InsertAtHead(10)
	dl.InsertAtHead(20)
	dl.InsertAtHead(30)

	dl.DisplayFromHead()
	dl.DisplayFromTail()

	fmt.Println("Head:", dl.Head.Data)
	fmt.Println("Tail:", dl.Tail.Data)
	fmt.Println("Head.Prev:", dl.Head.Prev.Data) // harus menunjuk ke Tail
	fmt.Println("Tail.Next:", dl.Tail.Next.Data) // harus menunjuk ke Head
}
