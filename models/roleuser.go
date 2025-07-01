package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique;not null"`
	Users []User `json:"users" gorm:"foreignKey:RoleID"`
}
