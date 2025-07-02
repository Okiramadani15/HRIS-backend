package controllers

import (
	"fmt"
	"hris-backend/config"
	"hris-backend/models"
	"hris-backend/utils"
	"os"
	"path/filepath"
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
		UserID      *uint   `json:"user_id,omitempty"`
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

	// Set UserID only if provided
	if input.UserID != nil && *input.UserID > 0 {
		employee.UserID = input.UserID
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
	PhotoURL    string  `json:"photo_url"`
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
	if strings.TrimSpace(input.PhotoURL) != "" {
		employee.PhotoURL = strings.TrimSpace(input.PhotoURL)
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

// UploadEmployeePhoto uploads employee photo
func UploadEmployeePhoto(c *fiber.Ctx) error {
	employeeID := c.Params("id")
	if strings.TrimSpace(employeeID) == "" {
		return utils.ValidationErrorResponse(c, "Employee ID is required")
	}

	// Check if employee exists
	var employee models.Employee
	if err := config.DB.First(&employee, employeeID).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee not found")
	}

	// Get uploaded file
	file, err := c.FormFile("photo")
	if err != nil {
		return utils.ValidationErrorResponse(c, "Photo file is required")
	}

	// Validate file type
	if !isValidImageType(file.Header.Get("Content-Type")) {
		return utils.ValidationErrorResponse(c, "Only JPG, JPEG, PNG files are allowed")
	}

	// Validate file size (max 5MB)
	if file.Size > 5*1024*1024 {
		return utils.ValidationErrorResponse(c, "File size must be less than 5MB")
	}

	// Create uploads directory
	uploadsDir := "./uploads/employees"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		return utils.InternalErrorResponse(c, "Failed to create upload directory")
	}

	// Generate unique filename
	filename := fmt.Sprintf("%s_%d%s", employee.EmployeeID, time.Now().Unix(), filepath.Ext(file.Filename))
	filePath := filepath.Join(uploadsDir, filename)

	// Save file
	if err := c.SaveFile(file, filePath); err != nil {
		return utils.InternalErrorResponse(c, "Failed to save photo")
	}

	// Update employee photo URL
	photoURL := fmt.Sprintf("/uploads/employees/%s", filename)
	employee.PhotoURL = photoURL
	if err := config.DB.Save(&employee).Error; err != nil {
		return utils.InternalErrorResponse(c, "Failed to update employee photo")
	}

	return utils.SuccessResponse(c, "Photo uploaded successfully", fiber.Map{
		"photo_url": photoURL,
		"employee": employee,
	})
}

// DeleteEmployeePhoto deletes employee photo
func DeleteEmployeePhoto(c *fiber.Ctx) error {
	employeeID := c.Params("id")
	if strings.TrimSpace(employeeID) == "" {
		return utils.ValidationErrorResponse(c, "Employee ID is required")
	}

	// Check if employee exists
	var employee models.Employee
	if err := config.DB.First(&employee, employeeID).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee not found")
	}

	// Delete physical file if exists
	if employee.PhotoURL != "" {
		filePath := fmt.Sprintf(".%s", employee.PhotoURL)
		if err := os.Remove(filePath); err != nil {
			// Log error but don't fail the request
			fmt.Printf("Warning: Failed to delete photo file: %v\n", err)
		}
	}

	// Update employee record
	employee.PhotoURL = ""
	if err := config.DB.Save(&employee).Error; err != nil {
		return utils.InternalErrorResponse(c, "Failed to update employee record")
	}

	return utils.SuccessResponse(c, "Photo deleted successfully", employee)
}

// Helper function to validate image types
func isValidImageType(contentType string) bool {
	validTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
	}
	for _, validType := range validTypes {
		if contentType == validType {
			return true
		}
	}
	return false
}

// LinkEmployeeToUser links an employee to a user account
func LinkEmployeeToUser(c *fiber.Ctx) error {
	employeeID := c.Params("id")
	
	type LinkRequest struct {
		UserEmail string `json:"user_email"`
	}
	
	var req LinkRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}
	
	if req.UserEmail == "" {
		return utils.ValidationErrorResponse(c, "User email is required")
	}
	
	// Find employee
	var employee models.Employee
	if err := config.DB.First(&employee, employeeID).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee not found")
	}
	
	// Find user by email
	var user models.User
	if err := config.DB.Where("email = ?", req.UserEmail).First(&user).Error; err != nil {
		return utils.NotFoundResponse(c, "User not found")
	}
	
	// Link employee to user
	employee.UserID = &user.ID
	if err := config.DB.Save(&employee).Error; err != nil {
		return utils.InternalErrorResponse(c, "Failed to link employee to user")
	}
	
	return utils.SuccessResponse(c, "Employee linked to user successfully", fiber.Map{
		"employee_id": employee.EmployeeID,
		"user_email":  user.Email,
		"user_name":   user.Name,
	})
}
