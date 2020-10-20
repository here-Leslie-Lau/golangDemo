package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

type Student struct {
	gorm.Model
	Student_name string `gorm:"not null;unique"`
	Grade        uint32 `gorm:"default:0"`
}

func (student Student) String() string {
	return "(student_name: " + student.Student_name + "\t-grade: " + strconv.Itoa(int(student.Grade))
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/gorm_demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("连接失败!")
	} else {
		fmt.Println("连接成功!")
	}

	defer db.Close()

	db.AutoMigrate(&Student{})

	//curd
	db.Create(&Student{Student_name: "leslie", Grade: 100})

	var student Student

	db.Where("student_name = ?", "leslie").First(&student)

	fmt.Println(student)

	db.Model(&student).Update("Student_name", "lion")

	fmt.Println(student)
}
