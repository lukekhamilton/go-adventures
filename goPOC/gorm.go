package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

import "fmt"

// Product ...
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Println("Hello gorm!")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect db")
	}
	defer db.Close()

	// migrate
	db.AutoMigrate(&Product{})

	//create
	db.Create(&Product{Code: "L1231231", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1231231")

	// update
	db.Model(&product).Update("Price", 2000)

	// delete
	db.Delete(&product)
}
