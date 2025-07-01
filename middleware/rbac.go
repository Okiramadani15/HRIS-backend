package middleware

import (
	"hris-backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// RequireRoles creates a middleware that checks user roles using dependency injection
func RequireRoles(db *gorm.DB, roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if db == nil {
			panic("Database connection cannot be nil in RBAC middleware")
		}
		
		// Safe type assertion to avoid panic
		userInterface := c.Locals("user")
		if userInterface == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		user, ok := userInterface.(*models.User)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid user context",
			})
		}

		var role models.Role
		if err := db.First(&role, user.RoleID).Error; err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Role not found",
			})
		}

		// Check if user has required role
		for _, r := range roles {
			if role.Name == r {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden - insufficient role permissions",
		})
	}
}
