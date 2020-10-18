package main

import "fmt"

//定义一个父结构体与两个子结构体
type Human struct {
	name    string
	age     int
	address string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	business string
}

//定义父结构体的一个method与重载Student的method
func (human Human) SayHi() {
	fmt.Println("hello,i'm Human")
}

func (student Student) SayHi() {
	fmt.Println("hello,i'm Student")
}

func main() {
	human := Human{"human", 18, "shenzhen"}
	human.SayHi()

	student := Student{Human{"student", 19, "shenzhen"}, "shenzhen university"}
	student.SayHi()

	employee := Employee{Human{"employee", 18, "shenzhen"}, "shenzhen business"}
	employee.SayHi()
}
