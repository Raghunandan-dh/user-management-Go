package models

import (
	"simple-REST/pkg/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (newUser *User) CreateUser() *User {
	db.Create(&newUser)
	return newUser
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(Id int64) *User {
	var user User
	db.Where("ID = ?", Id).Find(&user)
	return &user
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID = ?", ID).Delete(user)
	return user
}

func UpdateUser(user *User) {
	db.Save(&user)
}
