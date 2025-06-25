package models

import (
	"time"

	"gorm.io/gorm"
)

type LeaveRequest struct {
	gorm.Model
	EmployeeID  string    `json:"employee_id"`  // FK ke Employee.EmployeeID (string)
	LeaveTypeID uint      `json:"leave_type_id"` // FK ke LeaveType.ID (uint)
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status" gorm:"default:'pending'"` // pending, approved, rejected
	Note        string    `json:"note"`

	Employee  Employee  `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
	LeaveType LeaveType `json:"leave_type" gorm:"foreignKey:LeaveTypeID;references:ID"`
}
