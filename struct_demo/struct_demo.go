package main

import "fmt"

type Human struct {
	name    string
	age     int
	address string
}

type Student struct {
	Human
	grade int
}

func main() {
	human := Human{name: "leslie", age: 18, address: "shenzhen"}

	fmt.Println(human.name+" "+human.address+" ", human.age)

	student := Student{Human{name: "leslie", age: 18, address: "shenzhen"}, 100}

	fmt.Println(student.name+" "+student.address+" age:", student.age, " grade:", student.grade)
}
