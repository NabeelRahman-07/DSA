package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertFront(val int) {
	newNode := &Node{data: val}
	if l.head != nil {
		newNode.next = l.head
		l.head.prev = newNode
	}
	l.head = newNode
}

func (l *LinkedList) DeleteLast() {
	if l.head == nil{
		return
	}
	if l.head.next == nil{
		l.head = nil
		return
	}
	curr := l.head
	for curr.next != nil{
		curr = curr.next
	}
	curr.prev.next=nil
}

func(l *LinkedList)Display(){
	curr := l.head
	for curr != nil{
		fmt.Print(curr.data,"->")
		curr = curr.next
	}
	fmt.Println("nil")
}

func main(){
	list := LinkedList{}
	list.InsertFront(40)
	list.InsertFront(30)
	list.InsertFront(20)
	list.InsertFront(10)
	list.Display()
	list.DeleteLast()
	list.Display()
}
