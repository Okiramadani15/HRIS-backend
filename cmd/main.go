package main

import (
	"hris-backend/config"
	"hris-backend/models"
	"hris-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize database
	config.InitDB()
	
	// Migrate base tables first
	config.DB.AutoMigrate(
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
		&models.PayGrade{})
	
	// Migrate Employee without foreign key relationships
	config.DB.AutoMigrate(&models.Employee{})
	
	// Migrate Attendance last (has foreign key to Employee)
	config.DB.AutoMigrate(&models.Attendance{})

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Routes
	routes.SetupRoutes(app)

	// Start server
	log.Println("Server starting on :3000")
	log.Fatal(app.Listen(":3000"))
}
