package controllers

import (
	"hris-backend/config"
	"hris-backend/models"

	"github.com/gofiber/fiber/v2"
)

// CreateEducation
func CreateEducation(c *fiber.Ctx) error {
	education := new(models.Education)
	if err := c.BodyParser(education); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if education.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Education name is required"})
	}

	if err := config.DB.Create(education).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(education)
}

// Get All
func GetEducations(c *fiber.Ctx) error {
	var educations []models.Education
	config.DB.Find(&educations)
	return c.JSON(educations)
}

// Get by ID
func GetEducation(c *fiber.Ctx) error {
	id := c.Params("id")
	var education models.Education
	if err := config.DB.First(&education, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Education not found"})
	}
	return c.JSON(education)
}

// Update
func UpdateEducation(c *fiber.Ctx) error {
	id := c.Params("id")
	var education models.Education
	if err := config.DB.First(&education, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Education not found"})
	}

	var input models.Education
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	config.DB.Model(&education).Updates(input)
	return c.JSON(education)
}

// Delete
func DeleteEducation(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Education{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Education deleted successfully"})
}
