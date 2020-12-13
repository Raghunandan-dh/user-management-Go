package models

import (
	"simple-REST/pkg/config"

	"github.com/jinzhu/gorm"
)

type UserRoleMapping struct {
	gorm.Model
	UserId int64
	RoleId int64
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&UserRoleMapping{})
}

func (userRoleMapping *UserRoleMapping) CreateUserRole() {
	db.Create(&userRoleMapping)
}

func GetUserRoles(userId int64) []Role {
	var roles []Role
	db.Raw("SELECT * FROM roles r JOIN user_role_mappings urm ON urm.role_id=r.id WHERE urm.user_id= ?", userId).Scan(&roles)
	return roles
}
