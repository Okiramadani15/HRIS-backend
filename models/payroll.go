package models

import (
	"time"
	"gorm.io/gorm"
)

type Payroll struct {
	gorm.Model
	EmployeeID      string    `json:"employee_id" gorm:"type:varchar(50)"`
	Month           string    `json:"month"`
	Year            string    `json:"year"`
	BaseSalary      float64   `json:"base_salary"`
	TotalAllowances float64   `json:"total_allowances"`
	TotalDeductions float64   `json:"total_deductions"`
	NetSalary       float64   `json:"net_salary"`
	Status          string    `json:"status" gorm:"default:'generated'"`
	PaidAt          *time.Time `json:"paid_at,omitempty"`
}

type Allowance struct {
	gorm.Model
	EmployeeID string  `json:"employee_id" gorm:"type:varchar(50)"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
}

type Deduction struct {
	gorm.Model
	EmployeeID string  `json:"employee_id" gorm:"type:varchar(50)"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
}