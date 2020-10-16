package main

import "fmt"

func main() {
	number := 10

	switch number {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 10:
		fmt.Println("就是这儿!!!", number)
	default:
		fmt.Println("default")
	}
}
