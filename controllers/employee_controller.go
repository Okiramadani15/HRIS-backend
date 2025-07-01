package controllers

import (
	"fmt"
	"hris-backend/config"
	"hris-backend/models"
	"hris-backend/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// generateEmployeeID generates unique employee ID
func generateEmployeeID() string {
	// Get current count of employees
	var count int64
	config.DB.Model(&models.Employee{}).Count(&count)
	
	// Generate ID with format EMP + 4 digit number
	return fmt.Sprintf("EMP%04d", count+1)
}

// Create Employee
func CreateEmployee(c *fiber.Ctx) error {
	type EmployeeInput struct {
		UserID      uint    `json:"user_id"`
		EmployeeID  string  `json:"employee_id"` // Optional, will auto-generate if empty
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

	// Auto-generate Employee ID if not provided
	if strings.TrimSpace(input.EmployeeID) == "" {
		input.EmployeeID = generateEmployeeID()
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
	if input.FullName == "" || input.Position == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Full name and position are required"})
	}

	// Check if Employee ID already exists
	var existingEmployee models.Employee
	if err := config.DB.Where("employee_id = ?", strings.TrimSpace(input.EmployeeID)).First(&existingEmployee).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Employee ID already exists"})
	}

	employee := models.Employee{
		UserID:      input.UserID,
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

// UpdateInput represents input for updating employee
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

// UpdateEmployee updates employee with pointer optimization
func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	if strings.TrimSpace(id) == "" {
		panic("Employee ID parameter cannot be empty")
	}

	employee := &models.Employee{}
	if err := config.DB.First(employee, id).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee not found")
	}

	input := &UpdateInput{}
	if err := c.BodyParser(input); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	// Update fields if provided using pointer dereferencing
	if strings.TrimSpace(input.FullName) != "" {
		employee.FullName = strings.TrimSpace(input.FullName)
	}
	if strings.TrimSpace(input.Phone) != "" {
		employee.Phone = strings.TrimSpace(input.Phone)
	}
	if strings.TrimSpace(input.Address) != "" {
		employee.Address = strings.TrimSpace(input.Address)
	}
	if strings.TrimSpace(input.DateOfBirth) != "" {
		if dob, err := time.Parse("2006-01-02T15:04:05Z", input.DateOfBirth); err == nil {
			employee.DateOfBirth = dob
		} else if dob, err := time.Parse("2006-01-02", input.DateOfBirth); err == nil {
			employee.DateOfBirth = dob
		}
	}
	if strings.TrimSpace(input.Gender) != "" {
		employee.Gender = strings.TrimSpace(input.Gender)
	}
	if strings.TrimSpace(input.Position) != "" {
		employee.Position = strings.TrimSpace(input.Position)
	}
	if strings.TrimSpace(input.Department) != "" {
		employee.Department = strings.TrimSpace(input.Department)
	}
	if strings.TrimSpace(input.HireDate) != "" {
		if hireDate, err := time.Parse("2006-01-02T15:04:05Z", input.HireDate); err == nil {
			employee.HireDate = hireDate
		} else if hireDate, err := time.Parse("2006-01-02", input.HireDate); err == nil {
			employee.HireDate = hireDate
		}
	}
	if input.Salary > 0 {
		employee.Salary = input.Salary
	}
	if strings.TrimSpace(input.Status) != "" {
		employee.Status = strings.TrimSpace(input.Status)
	}

	if err := config.DB.Save(employee).Error; err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}

	// Load updated employee with user relationship
	config.DB.Preload("User").First(employee, employee.ID)

	return utils.SuccessResponse(c, "Employee updated successfully", employee)
}

// DeleteEmployee deletes employee with panic on critical errors
func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	if strings.TrimSpace(id) == "" {
		panic("Employee ID parameter cannot be empty")
	}

	employee := &models.Employee{}
	if err := config.DB.First(employee, id).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee not found")
	}

	// Check if employee has related records (attendances, leave requests)
	var attendanceCount, leaveRequestCount int64
	config.DB.Model(&models.Attendance{}).Where("employee_id = ?", employee.EmployeeID).Count(&attendanceCount)
	config.DB.Model(&models.LeaveRequest{}).Where("employee_id = ?", employee.EmployeeID).Count(&leaveRequestCount)
	
	if attendanceCount > 0 || leaveRequestCount > 0 {
		return utils.ValidationErrorResponse(c, "Cannot delete employee with existing attendance or leave records")
	}

	result := config.DB.Delete(employee)
	if result.Error != nil {
		return utils.InternalErrorResponse(c, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		panic("Employee not found for deletion")
	}

	return utils.SuccessResponse(c, "Employee deleted successfully", nil)
}

// GetMyProfile retrieves current user's employee profile
func GetMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return utils.ValidationErrorResponse(c, "User not authenticated")
	}

	employee := &models.Employee{}
	if err := config.DB.Preload("User").Where("user_id = ?", userID).First(employee).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee profile not found")
	}

	return utils.SuccessResponse(c, "Employee profile retrieved successfully", employee)
}
