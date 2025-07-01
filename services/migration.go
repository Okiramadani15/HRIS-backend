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

	// Migrate Payroll tables
	if err := db.AutoMigrate(
		&models.Payroll{},
		&models.Allowance{},
		&models.Deduction{},
	); err != nil {
		log.Fatal("Failed to migrate Payroll tables:", err)
	}

	// Migrate KPI tables
	if err := db.AutoMigrate(
		&models.KPICategory{},
		&models.KPIMetric{},
		&models.KPITarget{},
		&models.KPIActual{},
		&models.KPIReview{},
	); err != nil {
		log.Fatal("Failed to migrate KPI tables:", err)
	}

	// Seed default KPI categories
	seedKPICategories(db)

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

// seedKPICategories creates default KPI categories
func seedKPICategories(db *gorm.DB) {
	defaultCategories := []models.KPICategory{
		{Name: "Sales & Revenue", Description: "Sales performance and revenue generation metrics"},
		{Name: "Quality", Description: "Quality assurance and customer satisfaction metrics"},
		{Name: "Productivity", Description: "Work efficiency and output metrics"},
		{Name: "Customer Service", Description: "Customer service and support metrics"},
		{Name: "Innovation", Description: "Innovation and improvement initiatives"},
		{Name: "Leadership", Description: "Leadership and team management metrics"},
		{Name: "Learning & Development", Description: "Training and skill development metrics"},
	}

	for _, category := range defaultCategories {
		var count int64
		db.Model(&models.KPICategory{}).Where("name = ?", category.Name).Count(&count)
		if count == 0 {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create KPI category '%s': %v", category.Name, err)
			} else {
				log.Printf("Seeded KPI category '%s'", category.Name)
			}
		}
	}
}
