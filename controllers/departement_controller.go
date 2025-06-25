package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// CreateDepartment
func CreateDepartment(c *fiber.Ctx) error {
	var department models.Department

	if err := c.BodyParser(&department); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	department.Name = strings.TrimSpace(department.Name)
	department.Description = strings.TrimSpace(department.Description)

	if department.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Department name is required"})
	}

	if err := config.DB.Create(&department).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create department"})
	}

	return c.Status(fiber.StatusCreated).JSON(department)
}

// GetDepartments
func GetDepartments(c *fiber.Ctx) error {
	var departments []models.Department
	if err := config.DB.Find(&departments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch departments"})
	}
	return c.JSON(departments)
}

// GetDepartment by ID
func GetDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var department models.Department
	if err := config.DB.First(&department, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
	}
	return c.JSON(department)
}

// UpdateDepartment
func UpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var department models.Department

	if err := config.DB.First(&department, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
	}

	var input models.Department
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.Name != "" {
		department.Name = strings.TrimSpace(input.Name)
	}
	if input.Description != "" {
		department.Description = strings.TrimSpace(input.Description)
	}
	if input.Status != "" {
		department.Status = input.Status
	}

	if err := config.DB.Save(&department).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update department"})
	}

	return c.JSON(department)
}

// DeleteDepartment
func DeleteDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var department models.Department

	if err := config.DB.First(&department, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Department not found"})
	}

	if err := config.DB.Delete(&department).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete department"})
	}

	return c.JSON(fiber.Map{"message": "Department deleted successfully"})
}
