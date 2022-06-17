package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open("developer:1234567@/go_basics"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	user := User{
		FirstName: "Jonh",
		LastName:  "Doe",
		Email:     "john@doe.com",
	}

	db.Create(&user)
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
