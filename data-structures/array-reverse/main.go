package main

import "fmt"

func reverse(arr []int)[]int{
	first := 0
	last := len(arr)-1
	for first<last{
		arr[first], arr[last] = arr[last], arr[first]
		first++
		last--
	}
	return arr
}
func main(){
	arr := []int{1,2,3,4,5}
	fmt.Println(reverse(arr))
}