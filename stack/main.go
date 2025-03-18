package main

import (
	"errors"
	"fmt"

	"github.com/sony-nurdianto/GrokkingDataStructure/stack/linkedlist"
)

type Stack[T any] struct {
	data *linkedlist.SinglyLinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: linkedlist.NewSingleLinkedList[T](),
	}
}

func (s *Stack[T]) Len() int {
	return s.data.Len()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.data.Len() == 0
}

func (s *Stack[T]) Push(value T) {
	s.data.InsertAtHead(value)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("cannot pop empty stack")
	}

	return s.data.DeleteFromFront()
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("cannot peek empty stack")
	}

	return s.data.Head.Data, nil
}

func (s *Stack[T]) String() string {
	return s.data.String()
}

func main() {
	stack := NewStack[int]()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack:", stack)
	fmt.Println("Stack size:", stack.Len())

	val, _ := stack.Pop()
	fmt.Println("Popped value:", val)
	fmt.Println("Stack after pop:", stack)
}

// func main() {
// 	// Buat file profiling
// 	f, err := os.Create("cpu.prof")
// 	if err != nil {
// 		log.Fatal("Error creating CPU profile:", err)
// 	}
// 	pprof.StartCPUProfile(f)
// 	defer func() {
// 		pprof.StopCPUProfile()
// 		f.Close()
// 	}()
//
// 	// Lakukan operasi agar ada cukup aktivitas CPU
// 	stack := NewStack[int]()
// 	for i := range 1_000_000 {
// 		stack.Push(i)
// 	}
// 	for range 500_000 {
// 		stack.Pop()
// 	}
// }
