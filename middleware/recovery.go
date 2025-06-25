package middleware

import (
	"hris-backend/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

// RecoveryMiddleware handles panics and converts them to proper HTTP responses
func RecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered: %v", r)
				
				// Convert panic to proper error response
				switch v := r.(type) {
				case string:
					utils.InternalErrorResponse(c, v)
				case error:
					utils.InternalErrorResponse(c, v.Error())
				default:
					utils.InternalErrorResponse(c, "Internal server error")
				}
			}
		}()
		
		return c.Next()
	}
}