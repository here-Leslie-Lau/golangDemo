package main

import "fmt"

func main() {
	//遍历map集合
	myMap := map[string]string{"name": "leslie", "age": "18", "address": "shenzhen"}

	for key, value := range myMap {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}
