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

type CircularSinglyLinkedList[T any] struct {
	Head *Node[T]
}

// Menyisipkan node di awal
func (ll *CircularSinglyLinkedList[T]) InsertAtHead(data T) {
	newNode := &Node[T]{Data: data}

	if ll.Head == nil {
		newNode.Next = newNode // Circular reference
		ll.Head = newNode
		return
	}

	// Masukkan node di depan dan perbarui referensi terakhir
	tail := ll.Head
	for tail.Next != ll.Head {
		tail = tail.Next
	}

	newNode.Next = ll.Head
	tail.Next = newNode
	ll.Head = newNode
}

// Menyisipkan node setelah target tertentu
func (ll *CircularSinglyLinkedList[T]) InsertAfter(target, data T) {
	if ll.Head == nil {
		return
	}

	current := ll.Head
	for {
		if cmp.Equal(current.Data, target) {
			newNode := &Node[T]{Data: data, Next: current.Next}
			current.Next = newNode
			return
		}
		current = current.Next
		if current == ll.Head { // Jika kembali ke Head, berhenti
			break
		}
	}
}

// Menyisipkan node di akhir
func (ll *CircularSinglyLinkedList[T]) InsertAtTail(data T) {
	newNode := &Node[T]{Data: data}

	if ll.Head == nil {
		newNode.Next = newNode // Circular reference
		ll.Head = newNode
		return
	}

	tail := ll.Head
	for tail.Next != ll.Head {
		tail = tail.Next
	}

	tail.Next = newNode
	newNode.Next = ll.Head
}

// Menampilkan list secara melingkar
func (ll *CircularSinglyLinkedList[T]) Display() {
	if ll.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := ll.Head
	for {
		fmt.Print(current.Data, " -> ")
		current = current.Next
		if current == ll.Head {
			break
		}
	}
	fmt.Println("(circular back to Head)")
}

// Mencari node berdasarkan data
func (ll *CircularSinglyLinkedList[T]) Search(target T) *Node[T] {
	if ll.Head == nil {
		return nil
	}

	current := ll.Head
	for {
		if cmp.Equal(current.Data, target) {
			return current
		}
		current = current.Next
		if current == ll.Head {
			break
		}
	}
	return nil
}

// Menghapus node dengan data tertentu
func (ll *CircularSinglyLinkedList[T]) DeleteNode(target T) {
	if ll.Head == nil {
		return
	}

	// Jika node yang dihapus adalah Head
	if cmp.Equal(ll.Head.Data, target) {
		tail := ll.Head
		for tail.Next != ll.Head {
			tail = tail.Next
		}

		if ll.Head == tail { // Hanya satu node
			ll.Head = nil
		} else {
			tail.Next = ll.Head.Next
			ll.Head = ll.Head.Next
		}
		return
	}

	// Jika node yang dihapus bukan Head
	current := ll.Head
	for current.Next != ll.Head && !cmp.Equal(current.Next.Data, target) {
		current = current.Next
	}

	if current.Next != ll.Head { // Jika ditemukan
		current.Next = current.Next.Next
	}
}

func main() {
	ll := CircularSinglyLinkedList[int]{}

	ll.InsertAtHead(10)
	ll.InsertAtHead(20)
	ll.InsertAfter(20, 25)
	ll.InsertAtHead(30)
	ll.InsertAtTail(40)
	ll.InsertAtTail(50)
	ll.InsertAfter(40, 30)

	fmt.Println("Circular Linked List:")
	ll.Display() // Output: 30 -> 20 -> 25 -> 10 -> 40 -> 30 -> 50 -> (circular back to Head)

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
