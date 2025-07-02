package controllers

import (
	"hris-backend/models"
	"hris-backend/services"
	"hris-backend/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AttendanceController handles attendance operations with pointer optimization
type AttendanceController struct {
	dbService *services.DatabaseService
}

// NewAttendanceController creates a new attendance controller
func NewAttendanceController() *AttendanceController {
	return &AttendanceController{
		dbService: services.NewDatabaseService(),
	}
}

// Global instance for backward compatibility with lazy initialization
var attendanceController *AttendanceController

// getAttendanceController returns the controller instance with lazy initialization
func getAttendanceController() *AttendanceController {
	if attendanceController == nil {
		attendanceController = NewAttendanceController()
	}
	return attendanceController
}

// CreateAttendance creates a new attendance record with pointer optimization
func CreateAttendance(c *fiber.Ctx) error {
	return getAttendanceController().CreateAttendance(c)
}

func (ac *AttendanceController) CreateAttendance(c *fiber.Ctx) error {
	request := &models.AttendanceRequest{}
	
	if err := c.BodyParser(request); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	// Validate required fields
	if strings.TrimSpace(request.EmployeeID) == "" || strings.TrimSpace(request.Date) == "" || strings.TrimSpace(request.Status) == "" {
		return utils.ValidationErrorResponse(c, "Employee ID, date and status are required")
	}

	// Convert request to attendance model
	data, err := request.ToAttendance()
	if err != nil {
		return utils.ValidationErrorResponse(c, err.Error())
	}

	if err := ac.dbService.CreateAttendance(data); err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	
	return utils.SuccessResponse(c, "Attendance created successfully", data)
}

// GetAttendances retrieves all attendance records with pointer optimization
func GetAttendances(c *fiber.Ctx) error {
	return getAttendanceController().GetAttendances(c)
}

func (ac *AttendanceController) GetAttendances(c *fiber.Ctx) error {
	attendances := &[]models.Attendance{}
	
	// Get query parameters
	date := c.Query("date")
	employeeID := c.Query("employee_id")
	month := c.Query("month")
	
	if err := ac.dbService.GetAttendancesWithFilters(attendances, date, employeeID, month); err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	
	return utils.SuccessResponse(c, "Attendances retrieved successfully", *attendances)
}

// GetAttendanceByID retrieves attendance by ID with pointer optimization
func GetAttendanceByID(c *fiber.Ctx) error {
	return getAttendanceController().GetAttendanceByID(c)
}

func (ac *AttendanceController) GetAttendanceByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if strings.TrimSpace(idStr) == "" {
		panic("ID parameter cannot be empty")
	}
	
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.ValidationErrorResponse(c, "Invalid ID format")
	}
	
	attendance := &models.Attendance{} // Use pointer for efficiency
	if err := ac.dbService.GetAttendanceByID(uint(id), attendance); err != nil {
		return utils.NotFoundResponse(c, "Attendance not found")
	}
	
	return utils.SuccessResponse(c, "Attendance retrieved successfully", attendance)
}

// UpdateAttendance updates attendance record with pointer optimization
func UpdateAttendance(c *fiber.Ctx) error {
	return getAttendanceController().UpdateAttendance(c)
}

func (ac *AttendanceController) UpdateAttendance(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if strings.TrimSpace(idStr) == "" {
		panic("ID parameter cannot be empty")
	}
	
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.ValidationErrorResponse(c, "Invalid ID format")
	}
	
	attendance := &models.Attendance{}
	if err := ac.dbService.GetAttendanceByID(uint(id), attendance); err != nil {
		return utils.NotFoundResponse(c, "Attendance not found")
	}

	input := &models.AttendanceUpdateRequest{}
	if err := c.BodyParser(input); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	// Update only provided fields
	if strings.TrimSpace(input.Status) != "" {
		attendance.Status = input.Status
	}
	if strings.TrimSpace(input.Notes) != "" {
		attendance.Note = input.Notes
	}
	if input.OvertimeHours > 0 {
		attendance.OvertimeHours = input.OvertimeHours
	}

	if err := ac.dbService.UpdateAttendance(attendance); err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	
	return utils.SuccessResponse(c, "Attendance updated successfully", attendance)
}

// DeleteAttendance deletes attendance record with panic on critical errors
func DeleteAttendance(c *fiber.Ctx) error {
	return getAttendanceController().DeleteAttendance(c)
}

func (ac *AttendanceController) DeleteAttendance(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if strings.TrimSpace(idStr) == "" {
		panic("ID parameter cannot be empty")
	}
	
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.ValidationErrorResponse(c, "Invalid ID format")
	}
	
	if err := ac.dbService.DeleteAttendance(uint(id)); err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	
	return utils.SuccessResponse(c, "Attendance deleted successfully", nil)
}
