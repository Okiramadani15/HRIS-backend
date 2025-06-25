package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateLocation(c *fiber.Ctx) error {
	location := new(models.Location)
	if err := c.BodyParser(location); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	location.Name = strings.TrimSpace(location.Name)
	if location.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Location name is required"})
	}

	if err := config.DB.Create(&location).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(location)
}

func GetLocations(c *fiber.Ctx) error {
	var locations []models.Location
	config.DB.Find(&locations)
	return c.JSON(locations)
}

func GetLocation(c *fiber.Ctx) error {
	id := c.Params("id")
	var location models.Location
	if err := config.DB.First(&location, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Location not found"})
	}
	return c.JSON(location)
}

func UpdateLocation(c *fiber.Ctx) error {
	id := c.Params("id")
	var location models.Location
	if err := config.DB.First(&location, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Location not found"})
	}

	updateData := new(models.Location)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Model(&location).Updates(updateData)
	return c.JSON(location)
}

func DeleteLocation(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Location{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Location deleted"})
}
