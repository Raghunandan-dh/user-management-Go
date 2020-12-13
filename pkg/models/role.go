package models

import (
	"simple-REST/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Role struct {
	gorm.Model
	Name string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Role{})
}

func (newRole *Role) CreateRole() *Role {
	db.Create(&newRole)
	return newRole
}

func GetAllRoles() []Role {
	var Roles []Role
	db.Find(&Roles)
	return Roles
}

func GetRoleById(Id int64) *Role {
	var role Role
	db.Where("ID = ?", Id).Find(&role)
	return &role
}

func DeleteRole(ID int64) Role {
	var Role Role
	db.Where("ID = ?", ID).Delete(Role)
	return Role
}

func UpdateRole(Role *Role) {
	db.Save(&Role)
}
