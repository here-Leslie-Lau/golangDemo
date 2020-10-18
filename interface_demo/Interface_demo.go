package main

import "fmt"

//定义三结构体
type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

//定义一个interface,并通过Human实现,其中Student重写某一个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

//这里用指针传递是为了提升效率，避免内存过多对象
func (receiver *Human) SayHi() {
	fmt.Printf("hi,i'm a Human.My name is %s and my age is %d\n", receiver.name, receiver.age)
}

func (receiver Human) sing(lyrics string) {
	fmt.Printf("lalalaalalala,%s\n", lyrics)
}

func (receiver *Student) SayHi() {
	fmt.Printf("hi,i'm a Student.My name is %s and my age is %d\n", receiver.name, receiver.age)
}

func main() {

	human := Human{"human", 10}
	human.SayHi()
	human.sing("浮夸")

	student := Student{Human{"student", 10}, "shenzhen university"}
	student.SayHi()
	student.sing("学生之歌")

	employee := Employee{Human{"employee", 10}, "shenzhen company"}
	employee.SayHi()
	employee.sing("工作之歌")
}
