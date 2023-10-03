// main.go
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func GetUserByID(db *gorm.DB, userID uint) (User, error) {
	var user User
	result := db.First(&user, userID)
	return user, result.Error
}
