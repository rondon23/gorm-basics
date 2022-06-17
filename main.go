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

	////Create
	// user := User{
	// 	FirstName: "Jonh",
	// 	LastName:  "Doe",
	// 	Email:     "john@doe.com",
	// }

	// db.Create(&user)

	////Update
	// user := User{
	// 	Id:        1,
	// 	FirstName: "Jonh 2",
	// 	LastName:  "Doe 2",
	// 	Email:     "john@doe.com",
	// }

	// db.Updates(&user)

	//Delete
	// user := User{
	// 	Id:        1,
	// }

	// db.Delete(&user)

}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
