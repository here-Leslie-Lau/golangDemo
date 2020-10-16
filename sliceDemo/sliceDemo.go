package main

import "fmt"

func main() {
	array := [5]int {0,1,2,3,4}
	var slice []int = array[0:3]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println("cap=", cap(slice))
}