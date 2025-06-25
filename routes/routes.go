package routes

import (
	"hris-backend/config"
	"hris-backend/controllers"
	"hris-backend/middleware"
	"hris-backend/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public routes
	api := app.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// Protected routes
	protected := api.Group("/", middleware.JWTMiddleware)
	protected.Get("/profile", getProfile)

	// Employee routes
	protected.Post("/employees", controllers.CreateEmployee)
	protected.Get("/employees", controllers.GetEmployees)
	protected.Get("/employees/:id", controllers.GetEmployee)
	protected.Put("/employees/:id", controllers.UpdateEmployee)
	protected.Delete("/employees/:id", controllers.DeleteEmployee)
	protected.Get("/my-profile", controllers.GetMyProfile)

	// Position routes
	protected.Post("/positions", controllers.CreatePosition)
	protected.Get("/positions", controllers.GetAllPositions)
	protected.Get("/positions/:id", controllers.GetPositionByID)
	protected.Put("/positions/:id", controllers.UpdatePosition)
	protected.Delete("/positions/:id", controllers.DeletePosition)

	// Department routes
	protected.Post("/departments", controllers.CreateDepartment)
	protected.Get("/departments", controllers.GetDepartments)
	protected.Get("/departments/:id", controllers.GetDepartment)
	protected.Put("/departments/:id", controllers.UpdateDepartment)
	protected.Delete("/departments/:id", controllers.DeleteDepartment)

}

func getProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
