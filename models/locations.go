package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	Status   string `json:"status" gorm:"default:'active'"` // active/inactive
}
