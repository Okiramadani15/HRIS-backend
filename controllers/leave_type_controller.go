package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateLeaveType(c *fiber.Ctx) error {
	var leave models.LeaveType
	if err := c.BodyParser(&leave); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	leave.Name = strings.TrimSpace(leave.Name)
	if leave.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Leave type name is required"})
	}

	if err := config.DB.Create(&leave).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(leave)
}

func GetLeaveTypes(c *fiber.Ctx) error {
	var types []models.LeaveType
	if err := config.DB.Find(&types).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(types)
}

func GetLeaveType(c *fiber.Ctx) error {
	id := c.Params("id")
	var leave models.LeaveType
	if err := config.DB.First(&leave, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Leave type not found"})
	}
	return c.JSON(leave)
}

func UpdateLeaveType(c *fiber.Ctx) error {
	id := c.Params("id")
	var leave models.LeaveType
	if err := config.DB.First(&leave, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Leave type not found"})
	}

	var input models.LeaveType
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	leave.Name = strings.TrimSpace(input.Name)
	leave.Description = input.Description
	leave.Status = input.Status

	config.DB.Save(&leave)
	return c.JSON(leave)
}

func DeleteLeaveType(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.LeaveType{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Leave type deleted"})
}
