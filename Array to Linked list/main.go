package main

import (
	"fmt"
)

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func(l *LinkedList)ArrayToList(arr []int){
	if len(arr)==0{
		return 
	}
	l.head = &Node{data: arr[0]}
	temp := l.head
	for _,val := range arr{
		newNode:= &Node{data: val}
		temp.next = newNode
		temp = newNode
	}
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
	arr := []int{1,2,3,4,5,6}
	list := LinkedList{}
	list.ArrayToList(arr)
	list.Display()

}