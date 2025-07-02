package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	UserID       *uint     `json:"user_id,omitempty" gorm:"unique"`
	EmployeeID   string    `json:"employee_id" gorm:"unique"`
	FullName     string    `json:"full_name"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Gender       string    `json:"gender" gorm:"type:varchar(10)"`
	Position     string    `json:"position"`
	Department   string    `json:"department"`
	HireDate     time.Time `json:"hire_date"`
	Salary       float64   `json:"salary"`
	Status       string    `json:"status" gorm:"default:'active'"`
	PhotoURL     string    `json:"photo_url,omitempty"`
	
	// Relationship
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"` 
}