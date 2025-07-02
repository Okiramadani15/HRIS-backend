package middleware

import (
	"hris-backend/models"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// JWTMiddleware validates JWT token
func JWTMiddleware(c *fiber.Ctx) error {
	return JWTMiddlewareWithDB(nil)(c)
}

// JWTMiddlewareWithDB creates JWT middleware with database dependency injection
func JWTMiddlewareWithDB(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
	}

	// Handle both "Bearer token" and "token" formats
	var tokenString string
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		tokenString = authHeader
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := claims["user_id"]
			c.Locals("user_id", userID)
			
			// Load user with role for RBAC if database is provided
			if db != nil {
				var user models.User
				if err := db.Preload("Role").First(&user, userID).Error; err == nil {
					c.Locals("user", &user)
				}
			}
		}

		return c.Next()
	}
}
