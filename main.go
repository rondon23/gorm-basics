package main

import (
	"fmt"

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

	//Isert multiples
	// db.AutoMigrate(&User{})

	// users := []User{
	// 	{
	// 		FirstName: "John",
	// 		LastName:  "Doe",
	// 		Email:     "john@doe.com",
	// 	},
	// 	{
	// 		FirstName: "Jane",
	// 		LastName:  "Smith",
	// 		Email:     "jane@smith.com",
	// 	},
	// 	{
	// 		FirstName: "Willian",
	// 		LastName:  "Blake",
	// 		Email:     "willian@blake.com",
	// 	},
	// 	{
	// 		FirstName: "Bruno",
	// 		LastName:  "Rondon",
	// 		Email:     "bruno@rondon.com",
	// 	},
	// }
	// for _, user := range users {
	// 	db.Create(&user)
	// }

	// Queryng records
	// db.AutoMigrate(&User{})

	// user := User{}

	// db.First(&user)
	// db.Last(&user)
	// db.Where("last_name", "Smith").First(&user)

	fmt.Println(user)
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
