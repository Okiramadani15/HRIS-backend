package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreatePayGrade(c *fiber.Ctx) error {
	var grade models.PayGrade
	if err := c.BodyParser(&grade); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	grade.Name = strings.TrimSpace(grade.Name)
	grade.Description = strings.TrimSpace(grade.Description)

	if grade.Name == "" || grade.BaseSalary <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and base salary are required"})
	}

	if err := config.DB.Create(&grade).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(grade)
}

func GetPayGrades(c *fiber.Ctx) error {
	var grades []models.PayGrade
	if err := config.DB.Find(&grades).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(grades)
}

func GetPayGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	var grade models.PayGrade
	if err := config.DB.First(&grade, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pay grade not found"})
	}
	return c.JSON(grade)
}

func UpdatePayGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	var grade models.PayGrade
	if err := config.DB.First(&grade, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pay grade not found"})
	}

	var input models.PayGrade
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	grade.Name = strings.TrimSpace(input.Name)
	grade.Description = strings.TrimSpace(input.Description)
	grade.BaseSalary = input.BaseSalary
	grade.Status = input.Status

	config.DB.Save(&grade)
	return c.JSON(grade)
}

func DeletePayGrade(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.PayGrade{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Pay grade deleted"})
}
