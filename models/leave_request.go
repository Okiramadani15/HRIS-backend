package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LeaveRequest struct {
	gorm.Model
	EmployeeID       string    `json:"employee_id"`       // FK ke Employee.EmployeeID (string)
	LeaveTypeID      uint      `json:"leave_type_id"`     // FK ke LeaveType.ID (uint)
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	Reason           string    `json:"reason"`
	Status           string    `json:"status" gorm:"default:'pending'"` // pending, approved, rejected
	Note             string    `json:"note"`
	EmergencyContact string    `json:"emergency_contact"`

	Employee  Employee  `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
	LeaveType LeaveType `json:"leave_type" gorm:"foreignKey:LeaveTypeID;references:ID"`
}

// LeaveRequestRequest untuk parsing JSON request
type LeaveRequestRequest struct {
	EmployeeID       string `json:"employee_id"`
	LeaveTypeID      uint   `json:"leave_type_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	Reason           string `json:"reason"`
	EmergencyContact string `json:"emergency_contact,omitempty"`
}

// ToLeaveRequest converts LeaveRequestRequest to LeaveRequest model
func (lr *LeaveRequestRequest) ToLeaveRequest() (*LeaveRequest, error) {
	leaveRequest := &LeaveRequest{
		EmployeeID:       lr.EmployeeID,
		LeaveTypeID:      lr.LeaveTypeID,
		Reason:           lr.Reason,
		EmergencyContact: lr.EmergencyContact,
		Status:           "pending",
	}

	// Parse start date
	startDate, err := time.Parse("2006-01-02", lr.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start_date format: %v", err)
	}
	leaveRequest.StartDate = startDate

	// Parse end date
	endDate, err := time.Parse("2006-01-02", lr.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end_date format: %v", err)
	}
	leaveRequest.EndDate = endDate

	return leaveRequest, nil
}
