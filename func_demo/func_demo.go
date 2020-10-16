package main

import "fmt"

//函数练习

func sum(x, y int) int {
	return x + y
}

func main() {
	x, y := 1, 2

	sum := sum(x, y)

	fmt.Println("sum:", sum)
}
