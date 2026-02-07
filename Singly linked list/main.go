package main

import "fmt"

type Node struct{
	data int
	next *Node
}
type LinkedList struct{
	head *Node
}

func(l *LinkedList)InsertatBeginning(val int){
	newNode := &Node{data: val}
	newNode.next = l.head
	l.head = newNode
}

func(l *LinkedList)InsertAtEnd(val int){
	newNode := &Node{data : val}
	if l.head == nil {
		l.head = newNode
		return
	}
	temp := l.head
	for temp.next != nil{
		temp = temp.next
	}
	temp.next = newNode
}

func(l *LinkedList)Delete(val int){
	if l.head == nil{
		fmt.Println("List is empty")
		return
	}
	if l.head.data == val{
		l.head = l.head.next
		return
	}
	temp := l.head
	for temp.next != nil{
		if temp.next.data == val{
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
	fmt.Println("value not found")
}

func(l *LinkedList)Search(val int)bool{
	if l.head == nil {
		fmt.Println("Stack is empty")
		return false
	}
	temp := l.head
	for temp != nil{
		if temp.data == val {
			return true
		}
		temp = temp.next
	}
	return false
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
	list.InsertatBeginning(20)
	list.InsertatBeginning(10)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)
	list.Display()
	list.Delete(30)
	list.Display()
	fmt.Println(list.Search(90))
}