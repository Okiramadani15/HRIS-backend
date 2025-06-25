package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Create Employment Type
func CreateEmploymentType(c *fiber.Ctx) error {
	var input models.EmploymentType
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create employment type"})
	}

	return c.Status(fiber.StatusCreated).JSON(input)
}

// Get All Employment Types
func GetEmploymentTypes(c *fiber.Ctx) error {
	var types []models.EmploymentType
	if err := config.DB.Find(&types).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve employment types"})
	}
	return c.JSON(types)
}

// Get Employment Type by ID
func GetEmploymentType(c *fiber.Ctx) error {
	id := c.Params("id")
	var t models.EmploymentType
	if err := config.DB.First(&t, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employment type not found"})
	}
	return c.JSON(t)
}

// Update Employment Type
func UpdateEmploymentType(c *fiber.Ctx) error {
	id := c.Params("id")
	var t models.EmploymentType
	if err := config.DB.First(&t, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employment type not found"})
	}

	var input models.EmploymentType
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if input.Name != "" {
		t.Name = strings.TrimSpace(input.Name)
	}
	if input.Description != "" {
		t.Description = input.Description
	}
	if input.Status != "" {
		t.Status = input.Status
	}

	if err := config.DB.Save(&t).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update employment type"})
	}
	return c.JSON(t)
}

// Delete Employment Type
func DeleteEmploymentType(c *fiber.Ctx) error {
	id := c.Params("id")
	var t models.EmploymentType
	if err := config.DB.First(&t, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employment type not found"})
	}

	if err := config.DB.Delete(&t).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete employment type"})
	}
	return c.JSON(fiber.Map{"message": "Employment type deleted"})
}
