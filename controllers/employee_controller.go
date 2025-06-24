package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Create Employee
func CreateEmployee(c *fiber.Ctx) error {
	type EmployeeInput struct {
		EmployeeID  string  `json:"employee_id"`
		FullName    string  `json:"full_name"`
		Phone       string  `json:"phone"`
		Address     string  `json:"address"`
		DateOfBirth string  `json:"date_of_birth"`
		Gender      string  `json:"gender"`
		Position    string  `json:"position"`
		Department  string  `json:"department"`
		HireDate    string  `json:"hire_date"`
		Salary      float64 `json:"salary"`
	}

	var input EmployeeInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Get user ID from JWT
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	// Parse dates
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date of birth format (YYYY-MM-DD)"})
	}

	hireDate, err := time.Parse("2006-01-02", input.HireDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid hire date format (YYYY-MM-DD)"})
	}

	// Validate required fields
	if input.EmployeeID == "" || input.FullName == "" || input.Position == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Employee ID, full name, and position are required"})
	}

	employee := models.Employee{
		UserID:      uint(userID.(float64)),
		EmployeeID:  strings.TrimSpace(input.EmployeeID),
		FullName:    strings.TrimSpace(input.FullName),
		Phone:       strings.TrimSpace(input.Phone),
		Address:     strings.TrimSpace(input.Address),
		DateOfBirth: dob,
		Gender:      strings.TrimSpace(input.Gender),
		Position:    strings.TrimSpace(input.Position),
		Department:  strings.TrimSpace(input.Department),
		HireDate:    hireDate,
		Salary:      input.Salary,
		Status:      "active",
	}

	result := config.DB.Create(&employee)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Employee ID already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create employee"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Employee created successfully",
		"employee": employee,
	})
}

// Get All Employees
func GetEmployees(c *fiber.Ctx) error {
	var employees []models.Employee
	result := config.DB.Preload("User").Find(&employees)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch employees"})
	}

	return c.JSON(fiber.Map{
		"employees": employees,
		"total":     len(employees),
	})
}

// Get Employee by ID
func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee
	result := config.DB.Preload("User").First(&employee, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}

	return c.JSON(fiber.Map{"employee": employee})
}

// Update Employee
func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee
	result := config.DB.First(&employee, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}

	type UpdateInput struct {
		FullName    string  `json:"full_name"`
		Phone       string  `json:"phone"`
		Address     string  `json:"address"`
		DateOfBirth string  `json:"date_of_birth"`
		Gender      string  `json:"gender"`
		Position    string  `json:"position"`
		Department  string  `json:"department"`
		HireDate    string  `json:"hire_date"`
		Salary      float64 `json:"salary"`
		Status      string  `json:"status"`
	}

	var input UpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Update fields if provided
	if input.FullName != "" {
		employee.FullName = strings.TrimSpace(input.FullName)
	}
	if input.Phone != "" {
		employee.Phone = strings.TrimSpace(input.Phone)
	}
	if input.Address != "" {
		employee.Address = strings.TrimSpace(input.Address)
	}
	if input.DateOfBirth != "" {
		if dob, err := time.Parse("2006-01-02", input.DateOfBirth); err == nil {
			employee.DateOfBirth = dob
		}
	}
	if input.Gender != "" {
		employee.Gender = strings.TrimSpace(input.Gender)
	}
	if input.Position != "" {
		employee.Position = strings.TrimSpace(input.Position)
	}
	if input.Department != "" {
		employee.Department = strings.TrimSpace(input.Department)
	}
	if input.HireDate != "" {
		if hireDate, err := time.Parse("2006-01-02", input.HireDate); err == nil {
			employee.HireDate = hireDate
		}
	}
	if input.Salary > 0 {
		employee.Salary = input.Salary
	}
	if input.Status != "" {
		employee.Status = strings.TrimSpace(input.Status)
	}

	result = config.DB.Save(&employee)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update employee"})
	}

	return c.JSON(fiber.Map{
		"message":  "Employee updated successfully",
		"employee": employee,
	})
}

// Delete Employee
func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee
	result := config.DB.First(&employee, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}

	result = config.DB.Delete(&employee)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete employee"})
	}

	return c.JSON(fiber.Map{"message": "Employee deleted successfully"})
}

// Get My Employee Profile
func GetMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	var employee models.Employee
	result := config.DB.Preload("User").Where("user_id = ?", userID).First(&employee)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee profile not found"})
	}

	return c.JSON(fiber.Map{"employee": employee})
}
