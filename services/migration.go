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
		&models.Role{},
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

	// Seed default roles
	seedRoles(db)

	log.Println("Database migrations completed successfully")
}

// seedRoles creates default roles if they don't exist
func seedRoles(db *gorm.DB) {
	defaultRoles := []string{"admin", "hr", "employee"}

	for _, roleName := range defaultRoles {
		var count int64
		db.Model(&models.Role{}).Where("name = ?", roleName).Count(&count)
		if count == 0 {
			if err := db.Create(&models.Role{Name: roleName}).Error; err != nil {
				log.Printf("Failed to create role '%s': %v", roleName, err)
			} else {
				log.Printf("Seeded role '%s'", roleName)
			}
		}
	}
}
