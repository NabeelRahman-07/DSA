package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedLIst struct {
	head *Node
}

type Stack struct {
	top *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func InsertIntoArray(arr []int, index int, value int) []int {
	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = value
	return arr
}

func DeleteFromArray(arr []int, index int) []int {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func SearchFromArray(arr []int, val int) bool {
	for _, data := range arr {
		if data == val {
			return true
		}
	}
	return false
}

func (l *LinkedLIst) InsertAtBeginning(val int) {
	newNode := &Node{data: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedLIst) InsertAtEnd(val int) {
	newNode := &Node{data: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (l *LinkedLIst) Delete(val int) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		return
	}
	curr := l.head
	for curr.next != nil {
		if curr.next.data == val {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
	fmt.Println("Value not found")
}

func (l *LinkedLIst) ConvertToArray() []int {
	var result []int
	if l.head == nil {
		fmt.Println("List is empty")
		return result
	}
	curr := l.head
	for curr != nil {
		result = append(result, curr.data)
		curr = curr.next
	}
	return result
}

func (s *Stack) Push(val int) {
	newNode := &Node{data: val}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		fmt.Println("Stack is empty.")
		return 0, false
	}
	val := s.top.data
	s.top = s.top.next
	return val, true
}

func (s *Stack) Peek() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return 0
	}
	return s.top.data
}

func (q *Queue) Enqueue(val int) {
	newNode := &Node{data: val}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() (int, bool) {
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

func main() {
	
}
