package models

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	EmployeeID string     `json:"employee_id"`          // FK ke Employee.EmployeeID
	Date       time.Time  `json:"date" gorm:"not null"` // Tanggal absensi
	CheckIn    *time.Time `json:"check_in"`             // Jam masuk
	CheckOut   *time.Time `json:"check_out"`            // Jam keluar
	Status     string     `json:"status"`               // present, absent, leave, etc.
	Note       string     `json:"note"`                 // Keterangan

	Employee Employee `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
}
