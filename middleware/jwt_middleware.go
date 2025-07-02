package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"hris-backend/models"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// TokenBlacklist - Global blacklist untuk JWT tokens
var TokenBlacklist = make(map[string]time.Time)

// IsTokenBlacklisted - Check if token is blacklisted
func IsTokenBlacklisted(tokenString string) bool {
	tokenHash := hashToken(tokenString)
	
	expiryTime, exists := TokenBlacklist[tokenHash]
	if !exists {
		return false
	}

	// If token expiry has passed, remove from blacklist and return false
	if time.Now().After(expiryTime) {
		delete(TokenBlacklist, tokenHash)
		return false
	}

	return true
}

// HashToken - Export hash function untuk controller
func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func hashToken(token string) string {
	return HashToken(token)
}

// AddToBlacklist - Add token to blacklist
func AddToBlacklist(tokenString string, expiryTime time.Time) {
	tokenHash := hashToken(tokenString)
	TokenBlacklist[tokenHash] = expiryTime
}

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
	// Check if token is blacklisted
	if IsTokenBlacklisted(tokenString) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token has been revoked"})
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
