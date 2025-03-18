package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	data  []T
	front int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data:  make([]T, 0),
		front: 0,
	}
}

func (q *Queue[T]) Enqueue(data T) {
	q.data = append(q.data, data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.front >= len(q.data)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}

	value := q.data[q.front]
	q.front++

	if q.front > len(q.data)/2 {
		q.data = q.data[q.front:]
		q.front = 0
	}

	return value, nil
}

func (q *Queue[T]) PrintQueue() {
	fmt.Println(q.data[q.front:])
}

func (q *Queue[t]) PrintFrontIndex() {
	fmt.Println(q.front)
}

func main() {
	queue := NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.PrintFrontIndex()

	queue.PrintQueue()
}
