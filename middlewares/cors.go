package middlewares

import (
    "github.com/gofiber/fiber/v2"
)

func CorsMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Set CORS headers
        c.Set("Access-Control-Allow-Origin", "*")
        c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Set("Access-Control-Allow-Headers", "Content-Type, Origin, Authorization")

        // Handle preflight requests (OPTIONS)
        if c.Method() == fiber.MethodOptions {
            c.Status(fiber.StatusOK)
            return nil
        }

        
        // Continue to the next middleware or route handler
        return c.Next()
    }
}