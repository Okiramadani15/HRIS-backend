package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// CreatePosition - Tambah posisi baru
func CreatePosition(c *fiber.Ctx) error {
	var position models.Position

	if err := c.BodyParser(&position); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	position.Name = strings.TrimSpace(position.Name)

	if position.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Position name is required"})
	}

	// Cek duplikat
	var exists models.Position
	if err := config.DB.Where("name = ?", position.Name).First(&exists).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Position already exists"})
	}

	if err := config.DB.Create(&position).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create position"})
	}

	return c.Status(fiber.StatusCreated).JSON(position)
}

// GetAllPositions - Ambil semua data posisi
func GetAllPositions(c *fiber.Ctx) error {
	var positions []models.Position
	if err := config.DB.Find(&positions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch positions"})
	}
	return c.JSON(positions)
}

// GetPositionByID - Ambil posisi berdasarkan ID
func GetPositionByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var position models.Position
	if err := config.DB.First(&position, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
	}
	return c.JSON(position)
}

// UpdatePosition - Update data posisi
func UpdatePosition(c *fiber.Ctx) error {
	id := c.Params("id")
	var position models.Position

	if err := config.DB.First(&position, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
	}

	var updateData models.Position
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if updateData.Name != "" {
		updateData.Name = strings.TrimSpace(updateData.Name)
		position.Name = updateData.Name
	}

	if err := config.DB.Save(&position).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update position"})
	}

	return c.JSON(position)
}

// DeletePosition - Hapus posisi
func DeletePosition(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Position{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete position"})
	}
	return c.JSON(fiber.Map{"message": "Position deleted successfully"})
}
