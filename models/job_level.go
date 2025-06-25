package models

import "gorm.io/gorm"

type JobLevel struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:'active'"` // active / inactive
}
