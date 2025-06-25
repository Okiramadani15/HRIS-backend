package controllers

import (
	"hris-backend/config"
	"hris-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateMaritalStatus(c *fiber.Ctx) error {
	var status models.MaritalStatus
	if err := c.BodyParser(&status); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if status.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	if err := config.DB.Create(&status).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create marital status"})
	}
	return c.Status(fiber.StatusCreated).JSON(status)
}

func GetMaritalStatuses(c *fiber.Ctx) error {
	var statuses []models.MaritalStatus
	if err := config.DB.Find(&statuses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch statuses"})
	}
	return c.JSON(statuses)
}

func GetMaritalStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.MaritalStatus
	if err := config.DB.First(&status, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Status not found"})
	}
	return c.JSON(status)
}

func UpdateMaritalStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.MaritalStatus
	if err := config.DB.First(&status, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Status not found"})
	}

	var updateData models.MaritalStatus
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	status.Name = updateData.Name
	status.Status = updateData.Status

	config.DB.Save(&status)
	return c.JSON(status)
}

func DeleteMaritalStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.MaritalStatus{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete"})
	}
	return c.JSON(fiber.Map{"message": "Deleted successfully"})
}
