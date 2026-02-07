package main

import "fmt"

func findMax(arr []int)int{
	temp := arr[0]

	for _,val := range arr {
		if temp<val{
			temp = val
		}
	}
	return temp
}

func main(){
	a := []int{12,34,568,1,45,62,100,786,900,56,138,23}
	max := findMax(a)
	fmt.Println("Maximum value on the array:",max)   //no nested loops or repeated scanning. each elements checked once.
	fmt.Println("Time complexity: O(n)\nSpace complexity: O(1)")  //constant space.
}