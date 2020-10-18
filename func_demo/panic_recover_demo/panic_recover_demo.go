package main

import "fmt"

//抛出panic并用recover捕获
func testPanic() {
	fmt.Println("panic start...")
	defer func() {
		fmt.Println("defer start...")
		if temp := recover(); temp != nil {
			fmt.Println("recovered from...", temp)
		}
		fmt.Println("recover end...")
	}()
	panic("yep,i'm panic")
	fmt.Println("can you see me?")

}

func main() {
	fmt.Println("main start...")
	testPanic()
	fmt.Println("main end...")
}
