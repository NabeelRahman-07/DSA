package main

import "fmt"

func main(){
	val1 := new(int)
	val2 := new(string)
	myMap := make(map[string]int)
	sl := make([]int,5)
	fmt.Println(*val1)
	fmt.Println(*val2)
	fmt.Println(myMap)
	fmt.Println(sl)
}
