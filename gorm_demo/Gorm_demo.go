package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

type Product struct {
	gorm.Model
	Code  string
	Price float64
}

func (receiver Product) String() string {
	return "(code:" + receiver.Code + "-price:" + strconv.FormatFloat(receiver.Price, 'f', -1, 64) + ")"
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("connect wrong")
	}

	db.AutoMigrate(&Product{})

	//create
	db.Create(&Product{Code: "001", Price: 99.99})

	var product Product

	db.First(&product, "price = ?", "99.99")

	fmt.Println(product)

	db.Model(&product).Update("Price", 2000)

	fmt.Println(product)

	db.Delete(&product)

	defer db.Close()
}
