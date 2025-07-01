package models

import (
	"time"
	"gorm.io/gorm"
)

// KPICategory - Kategori KPI (Sales, Quality, Productivity, etc.)
type KPICategory struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100);unique"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

// KPIMetric - Template/Master KPI yang bisa digunakan
type KPIMetric struct {
	gorm.Model
	CategoryID    uint    `json:"category_id"`
	Name          string  `json:"name" gorm:"type:varchar(200)"`
	Description   string  `json:"description"`
	Unit          string  `json:"unit"` // %, number, currency, etc.
	TargetType    string  `json:"target_type"` // "higher_better", "lower_better", "exact"
	Weight        float64 `json:"weight" gorm:"default:1.0"` // Bobot untuk perhitungan total
	IsActive      bool    `json:"is_active" gorm:"default:true"`
	
	// Relationship
	Category KPICategory `json:"category" gorm:"foreignKey:CategoryID"`
}

// KPITarget - Target KPI untuk employee di periode tertentu
type KPITarget struct {
	gorm.Model
	EmployeeID   string    `json:"employee_id" gorm:"type:varchar(50)"`
	MetricID     uint      `json:"metric_id"`
	Period       string    `json:"period"` // "2025-Q1", "2025-07", "2025"
	PeriodType   string    `json:"period_type"` // "monthly", "quarterly", "yearly"
	TargetValue  float64   `json:"target_value"`
	MinValue     *float64  `json:"min_value,omitempty"`
	MaxValue     *float64  `json:"max_value,omitempty"`
	SetBy        string    `json:"set_by"` // User ID yang set target
	SetAt        time.Time `json:"set_at"`
	
	// Relationship
	Employee Employee  `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
	Metric   KPIMetric `json:"metric" gorm:"foreignKey:MetricID"`
}

// KPIActual - Pencapaian aktual KPI
type KPIActual struct {
	gorm.Model
	EmployeeID    string    `json:"employee_id" gorm:"type:varchar(50)"`
	MetricID      uint      `json:"metric_id"`
	Period        string    `json:"period"`
	PeriodType    string    `json:"period_type"`
	ActualValue   float64   `json:"actual_value"`
	Notes         string    `json:"notes"`
	RecordedBy    string    `json:"recorded_by"` // User ID yang input
	RecordedAt    time.Time `json:"recorded_at"`
	
	// Calculated fields
	TargetValue   *float64  `json:"target_value,omitempty" gorm:"-"`
	Achievement   *float64  `json:"achievement,omitempty" gorm:"-"` // Percentage
	Score         *float64  `json:"score,omitempty" gorm:"-"` // Weighted score
	
	// Relationship
	Employee Employee  `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
	Metric   KPIMetric `json:"metric" gorm:"foreignKey:MetricID"`
}

// KPIReview - Review/Evaluasi KPI periode tertentu
type KPIReview struct {
	gorm.Model
	EmployeeID     string    `json:"employee_id" gorm:"type:varchar(50)"`
	Period         string    `json:"period"`
	PeriodType     string    `json:"period_type"`
	TotalScore     float64   `json:"total_score"`
	MaxScore       float64   `json:"max_score"`
	Percentage     float64   `json:"percentage"`
	Grade          string    `json:"grade"` // A, B, C, D, E
	ReviewNotes    string    `json:"review_notes"`
	ReviewedBy     string    `json:"reviewed_by"`
	ReviewedAt     time.Time `json:"reviewed_at"`
	Status         string    `json:"status" gorm:"default:'draft'"` // draft, final, approved
	
	// Relationship
	Employee Employee `json:"employee" gorm:"foreignKey:EmployeeID;references:EmployeeID"`
}