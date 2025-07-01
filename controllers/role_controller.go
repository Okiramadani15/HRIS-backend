package controllers

import (
	"fmt"
	"hris-backend/config"
	"hris-backend/models"
	"hris-backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// generateUserID generates unique user ID
func generateUserID() string {
	// Get current count of users
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	
	// Generate ID with format USR + 4 digit number
	return fmt.Sprintf("USR%04d", count+1)
}

// CreateRole creates a new role (Admin only)
func CreateRole(c *fiber.Ctx) error {
	type RoleInput struct {
		Name string `json:"name"`
	}

	var input RoleInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role name is required"})
	}

	// Check if role already exists
	var existingRole models.Role
	if err := config.DB.Where("name = ?", strings.TrimSpace(input.Name)).First(&existingRole).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Role already exists"})
	}

	role := models.Role{
		Name: strings.TrimSpace(input.Name),
	}

	if err := config.DB.Create(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create role"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role created successfully",
		"role":    role,
	})
}

// GetRoles retrieves all roles
func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role
	if err := config.DB.Preload("Users").Find(&roles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch roles"})
	}

	return c.JSON(fiber.Map{
		"roles": roles,
		"total": len(roles),
	})
}

// GetRoleByID retrieves role by ID
func GetRoleByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var role models.Role
	if err := config.DB.Preload("Users").First(&role, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	return c.JSON(fiber.Map{"role": role})
}

// UpdateRole updates role information (Admin only)
func UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	type UpdateInput struct {
		Name string `json:"name"`
	}

	var input UpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if strings.TrimSpace(input.Name) != "" {
		// Check if new name already exists
		var existingRole models.Role
		if err := config.DB.Where("name = ? AND id != ?", strings.TrimSpace(input.Name), role.ID).First(&existingRole).Error; err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Role name already exists"})
		}
		role.Name = strings.TrimSpace(input.Name)
	}

	if err := config.DB.Save(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update role"})
	}

	return c.JSON(fiber.Map{
		"message": "Role updated successfully",
		"role":    role,
	})
}

// DeleteRole deletes a role (Admin only)
func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	// Check if role has users
	var userCount int64
	config.DB.Model(&models.User{}).Where("role_id = ?", id).Count(&userCount)
	if userCount > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot delete role that has users assigned"})
	}

	if err := config.DB.Delete(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete role"})
	}

	return c.JSON(fiber.Map{"message": "Role deleted successfully"})
}

// CreateUserWithRole creates user with specific role (Admin only)
func CreateUserWithRole(c *fiber.Ctx) error {
	type UserRoleInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   uint   `json:"role_id"`
	}

	var input UserRoleInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validate required fields
	if strings.TrimSpace(input.Name) == "" || strings.TrimSpace(input.Email) == "" || strings.TrimSpace(input.Password) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name, email, and password are required"})
	}

	if input.RoleID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role ID is required"})
	}

	// Check if role exists
	var role models.Role
	if err := config.DB.First(&role, input.RoleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	// Check if email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", strings.TrimSpace(input.Email)).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already exists"})
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		Name:     strings.TrimSpace(input.Name),
		Email:    strings.TrimSpace(input.Email),
		Password: hashedPassword,
		RoleID:   input.RoleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Load user with role
	config.DB.Preload("Role").First(&user, user.ID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role.Name,
		},
	})
}

// AssignUserRole assigns role to existing user (Admin only)
func AssignUserRole(c *fiber.Ctx) error {
	type AssignRoleInput struct {
		UserID uint `json:"user_id"`
		RoleID uint `json:"role_id"`
	}

	var input AssignRoleInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.UserID == 0 || input.RoleID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID and Role ID are required"})
	}

	// Check if user exists
	var user models.User
	if err := config.DB.First(&user, input.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Check if role exists
	var role models.Role
	if err := config.DB.First(&role, input.RoleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	// Update user role
	user.RoleID = input.RoleID
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to assign role"})
	}

	// Load updated user with role
	config.DB.Preload("Role").First(&user, input.UserID)

	return c.JSON(fiber.Map{
		"message": "Role assigned successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role.Name,
		},
	})
}