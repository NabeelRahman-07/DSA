package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func main() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
