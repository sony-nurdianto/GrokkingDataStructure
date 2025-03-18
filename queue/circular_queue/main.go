package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	data     []T
	capacity int
	front    int
	rear     int
	size     int
}

func NewCircularQueue[T any](cap int) *Queue[T] {
	return &Queue[T]{
		data:     make([]T, cap+1),
		capacity: cap + 1,
		front:    0,
		rear:     -1,
		size:     0,
	}
}

func (q *Queue[T]) IsFull() bool {
	return q.size == q.capacity-1
}

func (q *Queue[T]) Enqueue(data T) error {
	if q.IsEmpty() {
		return errors.New("queue is full")
	}

	q.rear = (q.rear + 1) % q.capacity
	// fmt.Println(q.rear)
	q.data[q.rear] = data
	q.size++

	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsFull() {
		return zero, errors.New("queue is full")
	}

	value := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	// fmt.Println(q.front)
	q.size--

	return value, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) PrintQueue() {
	fmt.Println(q.data)
}

func main() {
	queue := NewCircularQueue[int](5)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	// queue.Enqueue(5)

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	queue.Enqueue(5)
	queue.Enqueue(6)
	// if err := queue.Enqueue(5); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	queue.PrintQueue()
}
