package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

type User struct {
	//gorm.Model
	User_id       int
	User_name     string
	User_password string
}

func (receiver User) String() string {
	return "(user_id = " + strconv.Itoa(receiver.User_id) + "-user_name = " + receiver.User_name + "-user_password = " + receiver.User_password
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/gorm_demo?charset=utf8&parseTime=True&loc=Local")
	//db.SingularTable(true)

	if err != nil {
		fmt.Println(err)
		panic("connected wrong!")
	} else {
		fmt.Println("connected succeed!!!")
	}

	defer db.Close()

	//curd
	db.Create(&User{User_name: "leslie", User_password: "123"})

	var user User

	db.First(&user, "user_name = ?", "leslie")

	fmt.Println(user)

	db.Model(&user).Update("User_password", "123456")

	fmt.Println(user)

	db.Delete(&user)

	//var user User
	//fmt.Println(db.HasTable(user))
	//
	//var user_name="leslie"
	////条件查询
	//err = db.Where("user_name = ?", user_name).Find(&user).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(user.User_id)
}
