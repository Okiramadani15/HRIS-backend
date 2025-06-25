package services

import (
	"hris-backend/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running database migrations...")

	// Migrate base tables first
	if err := db.AutoMigrate(
		&models.User{},
		&models.Position{},
		&models.Department{},
		&models.Location{},
		&models.JobLevel{},
		&models.Education{},
		&models.MaritalStatus{},
		&models.Religion{},
		&models.EmploymentType{},
		&models.WorkShift{},
		&models.LeaveType{},
		&models.Bank{},
		&models.PayGrade{},
	); err != nil {
		log.Fatal("Failed to migrate base tables:", err)
	}

	// Migrate Employee table
	if err := db.AutoMigrate(&models.Employee{}); err != nil {
		log.Fatal("Failed to migrate Employee table:", err)
	}

	// Migrate tables with foreign keys
	if err := db.AutoMigrate(
		&models.Attendance{},
		&models.LeaveRequest{},
	); err != nil {
		log.Fatal("Failed to migrate tables with foreign keys:", err)
	}

	log.Println("Database migrations completed successfully")
}
