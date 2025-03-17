package main

import "cmp"

type Node[T cmp.Ordered] struct {
	Data T
	Next *Node[T]
}

func (n *Node[T]) GetData() T {
	return n.Data
}

func (n *Node[T]) NextData() *Node[T] {
	return n.Next
}

func (n *Node[T]) HasNext() bool {
	return n.Next != nil
}

func (n *Node[T]) Append(node *Node[T]) {
	n.Next = node
}

type SinglyLinkedList[T cmp.Ordered] struct {
	Head *Node[T]
}

func (ll *SinglyLinkedList[T]) InsertSortedList(data T) {
	newNode := &Node[T]{
		Data: data,
	}

	if ll.Head == nil || cmp.Less(data, ll.Head.Data) {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil && cmp.Less(current.Next.Data, data) {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
}

func main() {
}
