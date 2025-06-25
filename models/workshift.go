package models

import "gorm.io/gorm"

type WorkShift struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`                     // format: HH:MM
	EndTime     string `json:"end_time"`                       // format: HH:MM
	Status      string `json:"status" gorm:"default:'active'"` // active/inactive
}
