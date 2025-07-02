package controllers

import (
	"fmt"
	"hris-backend/config"
	"hris-backend/middleware"
	"hris-backend/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Menggunakan TokenBlacklist dari middleware

// Logout - Professional logout dengan token blacklist
func Logout(c *fiber.Ctx) error {
	// Get token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.ValidationErrorResponse(c, "Authorization header required")
	}

	// Extract token
	var tokenString string
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		tokenString = authHeader
	}

	if tokenString == "" {
		return utils.ValidationErrorResponse(c, "Token required")
	}

	// Parse token to get expiry time
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecret()), nil
	})

	var expiryTime time.Time
	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if exp, ok := claims["exp"].(float64); ok {
				expiryTime = time.Unix(int64(exp), 0)
			}
		}
	}

	// If we can't parse expiry, set a reasonable default (24 hours)
	if expiryTime.IsZero() {
		expiryTime = time.Now().Add(24 * time.Hour)
	}

	// Add to blacklist with expiry time
	middleware.AddToBlacklist(tokenString, expiryTime)

	// Log logout activity
	userID := c.Locals("user_id")
	logLogoutActivity(userID, c.IP())

	return utils.SuccessResponse(c, "Logged out successfully", fiber.Map{
		"logged_out_at": time.Now(),
		"message": "Token has been invalidated",
	})
}

// LogoutAll - Logout from all devices (invalidate all user tokens)
func LogoutAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return utils.ValidationErrorResponse(c, "User not authenticated")
	}

	// In a real implementation, you would:
	// 1. Store user token version in database
	// 2. Increment version to invalidate all tokens
	// 3. Check version in JWT middleware

	// For now, we'll just logout current token
	return Logout(c)
}

// Helper functions

func hashToken(token string) string {
	return middleware.HashToken(token)
}

func logLogoutActivity(userID interface{}, ipAddress string) {
	// Log to database or file
	fmt.Printf("[LOGOUT] User %v logged out from IP %s at %s\n", 
		userID, ipAddress, time.Now().Format("2006-01-02 15:04:05"))
}