package linkedlist

import (
	"fmt"
	"strings"

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
	size int // Menyimpan panjang linked list
}

func NewSingleLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Len mengembalikan jumlah elemen dalam linked list (O(1))
func (ll *SinglyLinkedList[T]) Len() int {
	return ll.size
}

func (ll *SinglyLinkedList[T]) InsertAtHead(data T) {
	newNode := &Node[T]{Data: data, Next: ll.Head}
	ll.Head = newNode
	ll.size++ // Tambah ukuran
}

func (ll *SinglyLinkedList[T]) InsertAfter(target, data T) {
	current := ll.Head

	for current != nil && !cmp.Equal(current.Data, target) {
		current = current.Next
	}

	if current != nil {
		newNode := &Node[T]{Data: data, Next: current.Next}
		current.Next = newNode
		ll.size++ // Tambah ukuran
	}
}

func (ll *SinglyLinkedList[T]) InsertAtTail(data T) {
	newNode := &Node[T]{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
		ll.size++
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.size++ // Tambah ukuran
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

func (ll *SinglyLinkedList[T]) DeleteFromFront() (T, error) {
	if ll.Head == nil {
		var zero T
		return zero, fmt.Errorf("list is empty")
	}

	removedData := ll.Head.Data
	ll.Head = ll.Head.Next
	ll.size-- // Kurangi ukuran

	return removedData, nil
}

func (ll *SinglyLinkedList[T]) DeleteNode(target T) {
	if ll.Head == nil {
		return
	}

	if cmp.Equal(ll.Head.Data, target) {
		ll.Head = ll.Head.Next
		ll.size-- // Kurangi ukuran
		return
	}

	current := ll.Head
	for current.Next != nil && !cmp.Equal(current.Next.Data, target) {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
		ll.size-- // Kurangi ukuran
	}
}

func (ll *SinglyLinkedList[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Head -> ")

	current := ll.Head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v -> ", current.Data))
		current = current.Next
	}

	sb.WriteString("nil")
	return sb.String()
}
