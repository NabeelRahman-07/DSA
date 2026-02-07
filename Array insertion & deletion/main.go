package main

import "fmt"

func insert(arr []int, index int, value int)[]int{
	arr = append(arr, 0)
	copy(arr[index+1:],arr[index:])
	arr[index]=value
	return arr
}
func delete(arr []int, index int)[]int{
	arr = append(arr[:index],arr[index+1:]... )
	return arr
}
func main(){
	arr := []int{10,20,30,40,50,60,70}
	fmt.Println(insert(arr, 2, 25))
	fmt.Println(delete(arr, 4))
}