package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	EmployeeID     string     `json:"employee_id"`          // FK ke Employee.EmployeeID
	Date           time.Time  `json:"date" gorm:"not null"` // Tanggal absensi
	CheckIn        *time.Time `json:"clock_in"`             // Jam masuk
	CheckOut       *time.Time `json:"clock_out"`            // Jam keluar
	Status         string     `json:"status"`               // present, absent, leave, etc.
	Note           string     `json:"notes"`                // Keterangan
	Location       string     `json:"location"`             // Lokasi absensi
	OvertimeHours  float64    `json:"overtime_hours"`       // Jam lembur

	Employee Employee `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
}

// AttendanceRequest untuk parsing JSON request
type AttendanceRequest struct {
	EmployeeID string `json:"employee_id"`
	Date       string `json:"date"`
	ClockIn    string `json:"clock_in,omitempty"`
	ClockOut   string `json:"clock_out,omitempty"`
	Status     string `json:"status"`
	Notes      string `json:"notes,omitempty"`
	Location   string `json:"location,omitempty"`
}

// AttendanceUpdateRequest untuk update operations
type AttendanceUpdateRequest struct {
	Status        string  `json:"status,omitempty"`
	Notes         string  `json:"notes,omitempty"`
	OvertimeHours float64 `json:"overtime_hours,omitempty"`
}

// ToAttendance converts AttendanceRequest to Attendance model
func (ar *AttendanceRequest) ToAttendance() (*Attendance, error) {
	attendance := &Attendance{
		EmployeeID: ar.EmployeeID,
		Status:     ar.Status,
		Note:       ar.Notes,
		Location:   ar.Location,
	}

	// Parse date
	date, err := time.Parse("2006-01-02", ar.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}
	attendance.Date = date

	// Parse clock_in if provided
	if ar.ClockIn != "" {
		clockIn, err := time.Parse("15:04:05", ar.ClockIn)
		if err != nil {
			return nil, fmt.Errorf("invalid clock_in format: %v", err)
		}
		// Combine date with time
		fullClockIn := time.Date(date.Year(), date.Month(), date.Day(), 
			clockIn.Hour(), clockIn.Minute(), clockIn.Second(), 0, time.Local)
		attendance.CheckIn = &fullClockIn
	}

	// Parse clock_out if provided
	if ar.ClockOut != "" {
		clockOut, err := time.Parse("15:04:05", ar.ClockOut)
		if err != nil {
			return nil, fmt.Errorf("invalid clock_out format: %v", err)
		}
		// Combine date with time
		fullClockOut := time.Date(date.Year(), date.Month(), date.Day(), 
			clockOut.Hour(), clockOut.Minute(), clockOut.Second(), 0, time.Local)
		attendance.CheckOut = &fullClockOut
	}

	return attendance, nil
}
