package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	v := make(chan bool)

	go func() {
		for {
			select {
			case o := <-c:
				fmt.Println(o)
			case <-time.After(3 * time.Second):
				fmt.Println("timeout")
				v <- true
				break
			}
		}
	}()
	<-v
}
