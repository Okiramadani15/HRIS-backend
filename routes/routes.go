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
