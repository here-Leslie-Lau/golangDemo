package main

import (
	"fmt"
	"reflect"
)

func main() {
	var price = 99.99

	value := reflect.ValueOf(&price)
	fmt.Println("the type of value is :", value.Type())
	fmt.Println("the kind of value is :", value.Kind())

	//改变值
	elem := value.Elem()
	elem.SetFloat(199.99)
	fmt.Println("the value of value after changing is:", price)
}
