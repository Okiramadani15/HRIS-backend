package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateWorkShift(c *fiber.Ctx) error {
	var shift models.WorkShift
	if err := c.BodyParser(&shift); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	shift.Name = strings.TrimSpace(shift.Name)

	if shift.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shift name is required"})
	}

	if err := config.DB.Create(&shift).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(shift)
}

func GetWorkShifts(c *fiber.Ctx) error {
	var shifts []models.WorkShift
	if err := config.DB.Find(&shifts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(shifts)
}

func GetWorkShift(c *fiber.Ctx) error {
	id := c.Params("id")
	var shift models.WorkShift
	if err := config.DB.First(&shift, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Work shift not found"})
	}
	return c.JSON(shift)
}

func UpdateWorkShift(c *fiber.Ctx) error {
	id := c.Params("id")
	var shift models.WorkShift
	if err := config.DB.First(&shift, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Work shift not found"})
	}

	var input models.WorkShift
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	shift.Name = strings.TrimSpace(input.Name)
	shift.Description = input.Description
	shift.StartTime = input.StartTime
	shift.EndTime = input.EndTime
	shift.Status = input.Status

	config.DB.Save(&shift)
	return c.JSON(shift)
}

func DeleteWorkShift(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.WorkShift{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Work shift deleted"})
}
