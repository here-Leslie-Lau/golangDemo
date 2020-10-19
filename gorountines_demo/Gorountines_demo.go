package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, num := range a {
		total += num
	}

	c <- total //send total to c
}

func main() {
	a := []int{0, 1, 2, 3, 4}
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
