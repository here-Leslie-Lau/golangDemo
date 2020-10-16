package main

import "fmt"

func main() {
	//统计数组的和
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("sum =", sum)
}
