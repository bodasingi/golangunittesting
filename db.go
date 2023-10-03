// main.go
package main

import (
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

type Database interface {
	First(dest interface{}, conds ...interface{}) *gorm.DB
}

func GetUserByID(db Database, userID uint) (User, error) {
	var user User
	result := db.First(&user, userID)
	return user, result.Error
}
