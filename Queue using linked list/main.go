package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type Queue struct{
	front *Node
	rear *Node
}

func (q *Queue)Enqueue(val int) {
	newNode := &Node{data: val}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue)Dequeue()(int, bool){
	if q.front == nil {
		fmt.Println("Queue is empty")
		return 0, false
	}
	val := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return val, true
}

func (q *Queue)DisplayQueue(){
	if q.front == nil {
		return
	}
	curr := q.front
	for curr != nil {
		fmt.Print(curr.data, "<-")
		curr = curr.next
	}
	fmt.Print("\n")
}

func main() {
	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.DisplayQueue()
	queue.Dequeue()
	queue.DisplayQueue()
}