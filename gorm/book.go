package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	db.AutoMigrate(Book{})

	// db.Create(Book{12,"Love"})

	// x := db.Where("id=?", 12)
	// fmt.Println(x)
	var b *Book
	db.First(&b, "id=?", 13)
	fmt.Println(b)
}
