package models

import "gorm.io/gorm"

type PayGrade struct {
	gorm.Model
	Name        string  `json:"name" gorm:"unique;not null"`    // e.g., "Grade A"
	Description string  `json:"description"`                    // e.g., "For senior-level employees"
	BaseSalary  float64 `json:"base_salary" gorm:"not null"`    // e.g., 12000000
	Status      string  `json:"status" gorm:"default:'active'"` // active/inactive
}
