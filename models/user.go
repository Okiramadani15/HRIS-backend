package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`

	RoleID uint `json:"role_id"`
	Role   Role `json:"role"`
}
