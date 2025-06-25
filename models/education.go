package models

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`    // Contoh: "S1"
	Description string `json:"description"`                    // Optional
	Status      string `json:"status" gorm:"default:'active'"` // active/inactive
}
