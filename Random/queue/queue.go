package main

import "fmt"

const maxCapacity = 10

type queue struct {
	arr               [maxCapacity]int
	rear, front, size int
}

func (q *queue) isFull() bool {
	return q.size == maxCapacity
}

func (q *queue) enqueue(element int) {
	if q.isFull() {
		fmt.Println("Queue is full")
		return
	}

	q.rear, q.size = q.rear+1, q.size+1
	q.arr[q.rear] = element
}

func (q *queue) display() {
	if q.rear+1 == q.front {
		fmt.Println("Queue is empty")
		return
	}

	for i := q.front; i <= q.rear; i++ {
		fmt.Println(q.arr[i])
	}
}

func (q *queue) dequeue() *int {
	if q.rear+1 == q.front {
		fmt.Println("Queue is empty")
		return nil
	}

	ele := q.arr[q.front]
	q.front++
	return &ele
}
func main() {
	q := queue{rear: -1, front: 0}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.display()
	q.dequeue()
	q.display()

}
