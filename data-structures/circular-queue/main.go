package main

import "fmt"

type Queue struct {
	arr   []int
	front int
	rear  int
	size  int
}

func (q *Queue) Enqueue(val int) bool {
	if (q.rear+1)%q.size == q.front {
		return false
	}
	if q.front == -1 {
		q.front = 0
		q.rear = 0
		q.arr[0] = val
		return true
	}
	q.rear = (q.rear + 1) % q.size
	q.arr[q.rear] = val
	return true
}

func (q *Queue) Dequeue() (int, bool) {
	if q.front == -1 {
		return 0, false
	}
	data := q.arr[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % q.size
	}
	return data, true
}
