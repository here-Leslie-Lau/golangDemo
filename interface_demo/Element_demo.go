package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

type Element interface{}

type list []Element

func (receiver Person) String() string {
	return "(name:" + receiver.name + "-(age:" + strconv.Itoa(receiver.age) + "years)"
}

func main() {
	list := make(list, 3)
	list[0] = 1                    //int
	list[1] = "Hello"              //string
	list[2] = Person{"leslie", 18} //Person

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("The value of list[%d] is %v.And the Type is %T\n", index, element, value)
		case string:
			fmt.Printf("The value of list[%d] is %v.And the Type is %T\n", index, element, value)
		case Person:
			fmt.Printf("The value of list[%d] is %v.And the Type is %T\n", index, element, value)
		default:
			fmt.Println("sorry.I don't know")
		}
	}
}
