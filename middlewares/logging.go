package middlewares

import (
    "github.com/gofiber/fiber/v2"
    "log"
)

func LoggingMiddleware(c *fiber.Ctx) error {
    log.Println(c.Method(), c.Path())
    return c.Next()
}
