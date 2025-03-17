package main

import (
	"cmp"
	"fmt"

	googlecmp "github.com/google/go-cmp/cmp"
)

type Node[T cmp.Ordered] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

type DoublyLinkedList[T cmp.Ordered] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (dl *DoublyLinkedList[T]) InsertAtHead(data T) {
	newNode := &Node[T]{Data: data}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode

		return
	}

	newNode.Next = dl.Head
	dl.Head.Prev = newNode
	dl.Head = newNode
}

func (dl *DoublyLinkedList[T]) InsertSorted(data T) {
	newNode := &Node[T]{Data: data}

	if dl.Head == nil {
		// List kosong, jadikan elemen pertama
		dl.Head = newNode
		dl.Tail = newNode
		return
	}

	if data <= dl.Head.Data {
		// Sisipkan di depan
		newNode.Next = dl.Head
		dl.Head.Prev = newNode
		dl.Head = newNode
		return
	}

	if data >= dl.Tail.Data {
		// Sisipkan di belakang
		newNode.Prev = dl.Tail
		dl.Tail.Next = newNode
		dl.Tail = newNode
		return
	}

	// Sisipkan di tengah
	current := dl.Head
	for current.Next != nil && current.Next.Data < data {
		current = current.Next
	}

	newNode.Next = current.Next
	newNode.Prev = current

	if current.Next != nil {
		current.Next.Prev = newNode
	}
	current.Next = newNode
}

// Helper function untuk menghindari nil pointer dereference
// func getNodeValue[T cmp.Ordered](node *Node[T]) any {
// 	if node == nil {
// 		return "nil"
// 	}
// 	return node.Data
// }
//
// // Fungsi debug untuk mencetak semua elemen dari Head ke Tail
// func printDebug[T cmp.Ordered](dl *DoublyLinkedList[T]) {
// 	fmt.Print("List: ")
// 	current := dl.Head
// 	for current != nil {
// 		fmt.Printf("[%v] ⇄ ", current.Data)
// 		current = current.Next
// 	}
// 	fmt.Println("nil")
// }

func (dl *DoublyLinkedList[T]) InsertAtMiddle(target, data T) {
	if dl.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := dl.Head
	for current != nil && !googlecmp.Equal(current.Data, target) {
		current = current.Next
	}

	if current == nil {
		return
	}

	newNode := &Node[T]{Data: data}
	newNode.Next = current.Next
	newNode.Prev = current

	if current.Next != nil {
		current.Next.Prev = newNode
	}

	current.Next = newNode

	if dl.Tail == current {
		dl.Tail = newNode
	}
}

func (dl *DoublyLinkedList[T]) InsertAtTail(data T) {
	newNode := &Node[T]{Data: data}

	if dl.Tail == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}

	dl.Tail.Next = newNode
	newNode.Prev = dl.Tail
	dl.Tail = newNode
}

func (dl *DoublyLinkedList[T]) DisplayFromHead() {
	if dl.Head == nil {
		fmt.Println("List is empty")
		return
	}

	// Cetak dari Head ke Tail
	fmt.Print("Forward: nil ⇄ ")
	current := dl.Head
	for current != nil {
		fmt.Print(current.Data, " ⇄ ")
		current = current.Next
	}
	fmt.Println("nil")
}

func (dl *DoublyLinkedList[T]) DisplayFromTail() {
	if dl.Tail == nil {
		fmt.Println("List is empty")
		return
	}

	fmt.Print("Backward: nil ⇄ ")
	current := dl.Tail
	for current != nil {
		fmt.Print(current.Data, " ⇄ ")
		current = current.Prev
	}
	fmt.Println("nil")
}

func (dl *DoublyLinkedList[T]) DeleteHead() {
	if dl.Head == nil {
		fmt.Println("List is empty")
		return
	}

	if dl.Head == dl.Tail {
		dl.Head = nil
		dl.Tail = nil
		return
	}

	dl.Head = dl.Head.Next
	dl.Head.Prev = nil
}

func (dl *DoublyLinkedList[T]) DeleteTail() {
	if dl.Tail == nil {
		fmt.Println("List is empty")
		return
	}

	if dl.Head == dl.Tail { // Jika hanya satu elemen
		dl.Head = nil
		dl.Tail = nil
		return
	}

	dl.Tail = dl.Tail.Prev
	dl.Tail.Next = nil
}

func (dl *DoublyLinkedList[T]) DeleteMiddle(target T) {
	if dl.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := dl.Head

	// Cari elemen yang ingin dihapus
	for current != nil && !googlecmp.Equal(current.Data, target) {
		current = current.Next
	}

	if current == nil {
		fmt.Println("Element not found")
		return
	}

	// Jika yang dihapus adalah Head
	if current == dl.Head {
		dl.DeleteHead()
		return
	}

	// Jika yang dihapus adalah Tail
	if current == dl.Tail {
		dl.DeleteTail()
		return
	}

	// Hapus node di tengah
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
}

func main() {
	dl := DoublyLinkedList[int]{}

	dl.InsertAtHead(10)
	dl.InsertAtHead(20)
	dl.InsertAtHead(30)
	// dl.InsertAtMiddle(10, 20)
	// dl.InsertAtTail(30)
	// dl.InsertAtTail(40)

	// dl.DisplayFromHead()
	// dl.DisplayFromTail()
	//
	fmt.Println(dl.Head.Data)
	fmt.Println(dl.Tail.Data)
}
