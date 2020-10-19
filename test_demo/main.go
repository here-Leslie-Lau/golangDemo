package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	add := Add(1, 2)
	fmt.Println("add is", add)
}
