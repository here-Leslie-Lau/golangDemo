package main

import "fmt"

func main() {
	i := 1

	goto Here

	fmt.Println("不走这里(反正看不到),i=", i+100000000000)

Here:
	fmt.Println("走这里,i=", i)
}
