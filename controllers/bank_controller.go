package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateBank(c *fiber.Ctx) error {
	var bank models.Bank
	if err := c.BodyParser(&bank); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	bank.Name = strings.TrimSpace(bank.Name)
	bank.Code = strings.TrimSpace(bank.Code)

	if bank.Name == "" || bank.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bank name and code are required"})
	}

	if err := config.DB.Create(&bank).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(bank)
}

func GetBanks(c *fiber.Ctx) error {
	var banks []models.Bank
	if err := config.DB.Find(&banks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(banks)
}

func GetBank(c *fiber.Ctx) error {
	id := c.Params("id")
	var bank models.Bank
	if err := config.DB.First(&bank, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bank not found"})
	}
	return c.JSON(bank)
}

func UpdateBank(c *fiber.Ctx) error {
	id := c.Params("id")
	var bank models.Bank
	if err := config.DB.First(&bank, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bank not found"})
	}

	var input models.Bank
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	bank.Name = strings.TrimSpace(input.Name)
	bank.Code = strings.TrimSpace(input.Code)
	bank.Status = input.Status

	config.DB.Save(&bank)
	return c.JSON(bank)
}

func DeleteBank(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Bank{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Bank deleted"})
}
