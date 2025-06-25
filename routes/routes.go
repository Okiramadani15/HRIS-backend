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

	// Location routes
	protected.Post("/locations", controllers.CreateLocation)
	protected.Get("/locations", controllers.GetLocations)
	protected.Get("/locations/:id", controllers.GetLocation)
	protected.Put("/locations/:id", controllers.UpdateLocation)
	protected.Delete("/locations/:id", controllers.DeleteLocation)

	// Job Level routes
	protected.Post("/job-levels", controllers.CreateJobLevel)
	protected.Get("/job-levels", controllers.GetJobLevels)
	protected.Get("/job-levels/:id", controllers.GetJobLevel)
	protected.Put("/job-levels/:id", controllers.UpdateJobLevel)
	protected.Delete("/job-levels/:id", controllers.DeleteJobLevel)

	// Education Routes
	protected.Post("/educations", controllers.CreateEducation)
	protected.Get("/educations", controllers.GetEducations)
	protected.Get("/educations/:id", controllers.GetEducation)
	protected.Put("/educations/:id", controllers.UpdateEducation)
	protected.Delete("/educations/:id", controllers.DeleteEducation)

	// Marital Status Routes
	protected.Post("/marital-statuses", controllers.CreateMaritalStatus)
	protected.Get("/marital-statuses", controllers.GetMaritalStatuses)
	protected.Get("/marital-statuses/:id", controllers.GetMaritalStatus)
	protected.Put("/marital-statuses/:id", controllers.UpdateMaritalStatus)
	protected.Delete("/marital-statuses/:id", controllers.DeleteMaritalStatus)

	//Religion Routes
	protected.Post("/religions", controllers.CreateReligion)
	protected.Get("/religions", controllers.GetReligions)
	protected.Get("/religions/:id", controllers.GetReligion)
	protected.Put("/religions/:id", controllers.UpdateReligion)
	protected.Delete("/religions/:id", controllers.DeleteReligion)

	//Employee Type Routes
	protected.Post("/employment-types", controllers.CreateEmploymentType)
	protected.Get("/employment-types", controllers.GetEmploymentTypes)
	protected.Get("/employment-types/:id", controllers.GetEmploymentType)
	protected.Put("/employment-types/:id", controllers.UpdateEmploymentType)
	protected.Delete("/employment-types/:id", controllers.DeleteEmploymentType)

	// Work Shift routes
	protected.Post("/work-shifts", controllers.CreateWorkShift)
	protected.Get("/work-shifts", controllers.GetWorkShifts)
	protected.Get("/work-shifts/:id", controllers.GetWorkShift)
	protected.Put("/work-shifts/:id", controllers.UpdateWorkShift)
	protected.Delete("/work-shifts/:id", controllers.DeleteWorkShift)

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
