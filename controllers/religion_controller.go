package controllers

import (
	"hris-backend/config"
	"hris-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateReligion(c *fiber.Ctx) error {
	var religion models.Religion
	if err := c.BodyParser(&religion); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if religion.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Religion name is required"})
	}

	if err := config.DB.Create(&religion).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create religion"})
	}
	return c.Status(fiber.StatusCreated).JSON(religion)
}

func GetReligions(c *fiber.Ctx) error {
	var religions []models.Religion
	if err := config.DB.Find(&religions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve data"})
	}
	return c.JSON(religions)
}

func GetReligion(c *fiber.Ctx) error {
	id := c.Params("id")
	var religion models.Religion
	if err := config.DB.First(&religion, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Religion not found"})
	}
	return c.JSON(religion)
}

func UpdateReligion(c *fiber.Ctx) error {
	id := c.Params("id")
	var religion models.Religion
	if err := config.DB.First(&religion, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Religion not found"})
	}

	var update models.Religion
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	religion.Name = update.Name
	religion.Status = update.Status

	config.DB.Save(&religion)
	return c.JSON(religion)
}

func DeleteReligion(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Religion{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete religion"})
	}
	return c.JSON(fiber.Map{"message": "Religion deleted successfully"})
}
