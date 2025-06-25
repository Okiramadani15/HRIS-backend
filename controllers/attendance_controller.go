package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"hris-backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateAttendance(c *fiber.Ctx) error {
	var data models.Attendance
	if err := c.BodyParser(&data); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	if strings.TrimSpace(data.EmployeeID) == "" || data.Date.IsZero() || strings.TrimSpace(data.Status) == "" {
		return utils.ValidationErrorResponse(c, "Employee ID, date and status are required")
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	return utils.SuccessResponse(c, "Attendance created successfully", data)
}

func GetAttendances(c *fiber.Ctx) error {
	var attendances []models.Attendance
	if err := config.DB.Preload("Employee").Find(&attendances).Error; err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	return utils.SuccessResponse(c, "Attendances retrieved successfully", attendances)
}

func GetAttendanceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var attendance models.Attendance
	if err := config.DB.Preload("Employee").First(&attendance, id).Error; err != nil {
		return utils.NotFoundResponse(c, "Attendance not found")
	}
	return utils.SuccessResponse(c, "Attendance retrieved successfully", attendance)
}

func UpdateAttendance(c *fiber.Ctx) error {
	id := c.Params("id")
	var attendance models.Attendance
	if err := config.DB.First(&attendance, id).Error; err != nil {
		return utils.NotFoundResponse(c, "Attendance not found")
	}

	var input models.Attendance
	if err := c.BodyParser(&input); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	attendance.CheckIn = input.CheckIn
	attendance.CheckOut = input.CheckOut
	attendance.Status = input.Status
	attendance.Note = input.Note

	config.DB.Save(&attendance)
	return utils.SuccessResponse(c, "Attendance updated successfully", attendance)
}

func DeleteAttendance(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Attendance{}, id).Error; err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	return utils.SuccessResponse(c, "Attendance deleted successfully", nil)
}
