package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateJobLevel(c *fiber.Ctx) error {
	var jobLevel models.JobLevel
	if err := c.BodyParser(&jobLevel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	jobLevel.Name = strings.TrimSpace(jobLevel.Name)
	if jobLevel.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Job level name is required"})
	}

	if err := config.DB.Create(&jobLevel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create job level"})
	}

	return c.Status(fiber.StatusCreated).JSON(jobLevel)
}

func GetJobLevels(c *fiber.Ctx) error {
	var jobLevels []models.JobLevel
	if err := config.DB.Find(&jobLevels).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve job levels"})
	}
	return c.JSON(jobLevels)
}

func GetJobLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	var jobLevel models.JobLevel
	if err := config.DB.First(&jobLevel, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job level not found"})
	}
	return c.JSON(jobLevel)
}

func UpdateJobLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	var jobLevel models.JobLevel
	if err := config.DB.First(&jobLevel, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job level not found"})
	}

	var update models.JobLevel
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if update.Name != "" {
		jobLevel.Name = update.Name
	}
	if update.Description != "" {
		jobLevel.Description = update.Description
	}
	if update.Status != "" {
		jobLevel.Status = update.Status
	}

	if err := config.DB.Save(&jobLevel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update job level"})
	}

	return c.JSON(jobLevel)
}

func DeleteJobLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.JobLevel{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete job level"})
	}
	return c.JSON(fiber.Map{"message": "Job level deleted successfully"})
}
