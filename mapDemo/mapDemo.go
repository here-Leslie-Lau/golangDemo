package main

import "fmt"

func main()  {
	myMap := map[string]string {"name" : "leslie", "age" : "18", "address" : "shenzhen", "lalala" : "yep"}

	//fmt.Println(myMap["name"])
	//fmt.Println(myMap["age"])
	//fmt.Println(myMap["address"])
	//fmt.Println("length:", len(myMap))
	//
	result, ok := myMap["lalala"]

	if ok {
		fmt.Println("myMap[\"lalala\"]=", result)
	} else {
		fmt.Println("sorry,myMap have no result")
	}
}