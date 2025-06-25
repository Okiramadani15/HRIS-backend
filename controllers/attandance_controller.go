package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateAttendance(c *fiber.Ctx) error {
	var data models.Attendance
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if strings.TrimSpace(data.EmployeeID) == "" || data.Date.IsZero() || strings.TrimSpace(data.Status) == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Employee ID, date and status are required"})
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(data)
}

func GetAttendances(c *fiber.Ctx) error {
	var attendances []models.Attendance
	if err := config.DB.Preload("Employee").Find(&attendances).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(attendances)
}

func GetAttendanceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var attendance models.Attendance
	if err := config.DB.Preload("Employee").First(&attendance, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Attendance not found"})
	}
	return c.JSON(attendance)
}

func UpdateAttendance(c *fiber.Ctx) error {
	id := c.Params("id")
	var attendance models.Attendance
	if err := config.DB.First(&attendance, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}

	var input models.Attendance
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	attendance.CheckIn = input.CheckIn
	attendance.CheckOut = input.CheckOut
	attendance.Status = input.Status
	attendance.Note = input.Note

	config.DB.Save(&attendance)
	return c.JSON(attendance)
}

func DeleteAttendance(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Attendance{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Attendance deleted"})
}
