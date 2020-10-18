package main

import "fmt"

//defer会在函数返回前执行
//defer底层是个栈，多个defer时，会采用先进后出的方式
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("i=", i)
	}
}
