package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique;not null"`
	Code   string `json:"code" gorm:"unique;not null"` // e.g. "014" untuk BCA
	Status string `json:"status" gorm:"default:'active'"`
}
