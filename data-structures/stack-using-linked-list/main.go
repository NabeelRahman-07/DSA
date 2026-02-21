package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(val int) {
	newNode := &Node{data: val}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) Pop() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return -1
	}
	val := s.top.data
	s.top = s.top.next
	return val
}

func (s *Stack)DisplayStack(){
	if s.top == nil {
		return
	}
	curr := s.top
	for curr != nil {
		fmt.Println("|",curr.data,"|")
		curr = curr.next
	}
	fmt.Print("\n")
}

func main() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.DisplayStack()
	stack.Pop()
	stack.DisplayStack()
	
}
