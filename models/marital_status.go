package models

import "gorm.io/gorm"

type MaritalStatus struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique;not null"`
	Status string `json:"status" gorm:"default:'active'"` // active/inactive
}
