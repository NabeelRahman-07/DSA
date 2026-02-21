package main

import "fmt"

type Queue struct{
	items []int
}

func (q *Queue)Enqueue(val int){
	q.items = append(q.items, val)
}

func (q *Queue)Dequeue()(int, bool) {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return 0, false
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

func main(){
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println(queue.items)
	queue.Dequeue()
	fmt.Println(queue.items)
}